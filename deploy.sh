#!/bin/bash

# Deploy script for Go API on Heroku
# This script automatically increments version, builds, commits, and deploys

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to get the latest version tag
get_latest_version() {
    local latest_tag=$(git tag --sort=-version:refname | head -1)
    if [[ -z "$latest_tag" ]]; then
        echo "v1.00.00"
    else
        echo "$latest_tag"
    fi
}

# Function to increment version based on commit type
increment_version() {
    local version=$1
    local commit_type=${2:-other}  # Default to other
    
    # Remove 'v' prefix if present
    version=${version#v}
    
    # Split version into parts
    IFS='.' read -ra VERSION_PARTS <<< "$version"
    local major=${VERSION_PARTS[0]:-1}
    local minor=${VERSION_PARTS[1]:-0}
    local patch=${VERSION_PARTS[2]:-0}
    
    case $commit_type in
        feat)
            # Feature commits increment minor version and reset patch to 00
            minor=$((minor + 1))
            patch=0
            ;;
        other)
            # Other commits increment patch version
            patch=$((patch + 1))
            ;;
        major)
            # Manual major increment
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        minor)
            # Manual minor increment
            minor=$((minor + 1))
            patch=0
            ;;
        patch)
            # Manual patch increment
            patch=$((patch + 1))
            ;;
        *)
            print_error "Invalid increment type: $commit_type. Use feat, other, major, minor, or patch"
            exit 1
            ;;
    esac
    
    echo "v${major}.$(printf "%02d" $minor).$(printf "%02d" $patch)"
}

# Function to check if we're on main/master branch
check_branch() {
    local current_branch=$(git branch --show-current)
    if [[ "$current_branch" != "master" && "$current_branch" != "main" ]]; then
        print_warning "You're currently on branch: $current_branch"
        print_warning "Consider switching to main/master branch for deployment"
        read -p "Do you want to continue? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            print_error "Deployment cancelled"
            exit 1
        fi
    fi
}

# Function to check for uncommitted changes
check_working_tree() {
    if ! git diff-index --quiet HEAD --; then
        print_warning "You have uncommitted changes in your working tree"
        git status --short
        read -p "Do you want to commit these changes first? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            read -p "Enter commit message: " commit_msg
            if [[ -z "$commit_msg" ]]; then
                commit_msg="chore: commit changes before deployment"
            fi
            git add .
            git commit -m "$commit_msg"
            print_success "Changes committed"
        else
            print_error "Please commit or stash your changes before deploying"
            exit 1
        fi
    fi
}

# Function to build the Go binary
build_binary() {
    print_status "Building Go binary..."
    
    # Clean previous builds
    if [[ -f "main" ]]; then
        rm main
        print_status "Removed previous binary"
    fi
    
    # Build the binary
    if go build -o main; then
        print_success "Binary built successfully"
    else
        print_error "Failed to build binary"
        exit 1
    fi
}

# Function to commit and tag the new version
commit_version() {
    local new_version=$1
    
    print_status "Committing new version: $new_version"
    
    # Add the binary and any other changes
    git add main
    git add . || true
    
    # Commit with version
    git commit -m "chore(release): $new_version"
    
    # Create and push the tag
    git tag "$new_version"
    
    print_success "Version $new_version committed and tagged"
}

# Function to deploy to Heroku
deploy_to_heroku() {
    print_status "Deploying to Heroku..."
    
    # Check if Heroku remote exists
    if ! git remote | grep -q heroku; then
        print_error "Heroku remote not found. Please add it with:"
        print_error "heroku git:remote -a YOUR_APP_NAME"
        exit 1
    fi
    
    # Push to Heroku
    if git push heroku master; then
        print_success "Successfully deployed to Heroku!"
        print_status "Your app should be available at: https://$(heroku info -s | grep web_url | cut -d= -f2)"
    else
        print_error "Failed to deploy to Heroku"
        exit 1
    fi
}

# Function to push to origin
push_to_origin() {
    print_status "Pushing to origin..."
    
    if git push origin master; then
        print_success "Pushed to origin successfully"
    else
        print_warning "Failed to push to origin"
    fi
    
    if git push origin --tags; then
        print_success "Pushed tags to origin successfully"
    else
        print_warning "Failed to push tags to origin"
    fi
}

# Main deployment function
main() {
    print_status "Starting deployment process..."
    
    # Check if we're in a git repository
    if ! git rev-parse --git-dir > /dev/null 2>&1; then
        print_error "Not in a git repository"
        exit 1
    fi
    
    # Check branch and working tree
    check_branch
    check_working_tree
    
    # Get current version and determine increment type
    local current_version=$(get_latest_version)
    local increment_type=${1:-auto}
    
    # Auto-detect commit type if not specified
    if [[ "$increment_type" == "auto" ]]; then
        # Check if there are any feat commits since last tag
        local last_tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
        if [[ -n "$last_tag" ]]; then
            # Look for various feat commit patterns: feat:, feat(), feat, etc.
            local feat_commits=$(git log --oneline "$last_tag"..HEAD | grep -E -c "(feat:|feat\(|feat\s|^feat)" || echo "0")
            if [[ $feat_commits -gt 0 ]]; then
                increment_type="feat"
                print_status "Detected $feat_commits feature commit(s) - will increment minor version"
            else
                increment_type="other"
                print_status "No feature commits detected - will increment patch version"
            fi
        else
            increment_type="other"
        fi
    fi
    
    local new_version=$(increment_version "$current_version" "$increment_type")
    
    print_status "Current version: $current_version"
    print_status "New version: $new_version"
    print_status "Increment type: $increment_type"
    
    # Confirm deployment
    read -p "Proceed with deployment? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_error "Deployment cancelled"
        exit 1
    fi
    
    # Execute deployment steps
    build_binary
    commit_version "$new_version"
    deploy_to_heroku
    push_to_origin
    
    print_success "Deployment completed successfully!"
    print_status "Version $new_version is now live on Heroku"
}

# Show usage if help is requested
if [[ "$1" == "-h" || "$1" == "--help" ]]; then
    echo "Usage: $0 [increment_type]"
    echo ""
    echo "Increment types:"
    echo "  auto    - Auto-detect based on commit types (default)"
    echo "            feat: commits → minor version increment"
    echo "            other commits → patch version increment"
    echo "  feat    - Force feature increment (minor version)"
    echo "  other   - Force other increment (patch version)"
    echo "  major   - Manual major version increment"
    echo "  minor   - Manual minor version increment"
    echo "  patch   - Manual patch version increment"
    echo ""
    echo "Examples:"
    echo "  $0           # Auto-detect commit type and increment"
    echo "  $0 feat      # Force feature increment (minor)"
    echo "  $0 other     # Force other increment (patch)"
    echo "  $0 major     # Manual major increment"
    echo ""
    echo "This script will:"
    echo "  1. Check git status and branch"
    echo "  2. Auto-detect commit types and increment version"
    echo "  3. Build Go binary"
    echo "  4. Commit and tag new version"
    echo "  5. Deploy to Heroku"
    echo "  6. Push to origin"
    
    exit 0
fi

# Run main function with all arguments
main "$@"
