# Deployment Guide

This guide explains how to use the automated deployment scripts for your Go API on Heroku.

## 🚀 Available Scripts

### 1. `deploy.sh` - Full Featured Deployment Script

**Features:**
- Automatic version incrementing (patch, minor, or major)
- Git status checking and validation
- Interactive prompts for safety
- Colored output for better readability
- Comprehensive error handling
- Branch validation

**Usage:**
```bash
# Auto-detect commit types and increment (default)
./deploy.sh

# Force feature increment (minor version)
./deploy.sh feat

# Force other increment (patch version)
./deploy.sh other

# Manual version increments
./deploy.sh major
./deploy.sh minor
./deploy.sh patch

# Show help
./deploy.sh --help
```

**What it does:**
1. ✅ Checks git status and branch
2. 🔍 Validates working tree
3. 📈 Automatically increments version based on latest git tag
4. 🔨 Builds Go binary
5. 💾 Commits and tags new version
6. 🚀 Deploys to Heroku
7. 📤 Pushes to origin repository

### 2. `quick-deploy.sh` - Simple Quick Deployment

**Features:**
- Always increments patch version
- No interactive prompts
- Fast deployment for routine updates
- Emoji-based status indicators

**Usage:**
```bash
./quick-deploy.sh
```

**What it does:**
1. 🔍 Gets latest version tag
2. 📈 Increments patch version
3. 🔨 Builds Go binary
4. 💾 Commits and tags
5. 🚀 Deploys to Heroku
6. 📤 Pushes to origin

## 📋 Prerequisites

Before using these scripts, ensure you have:

1. **Go installed** and accessible via `go` command
2. **Git repository** initialized and configured
3. **Heroku CLI** installed and authenticated
4. **Heroku remote** added to your git repository:
   ```bash
   heroku git:remote -a YOUR_APP_NAME
   ```

## 🔧 Smart Version Management

The scripts automatically manage semantic versioning based on your commit types:

- **Current version**: `v1.02.56`
- **Feature commits** (`feat:`) → `v1.03.00` (minor increment, patch resets to 00)
- **Other commits** → `v1.02.57` (patch increment)
- **Manual control** available for major/minor/patch increments

**How it works:**
1. **Auto-detection**: Scripts scan commits since last tag for feature commit patterns
2. **Feature commit patterns detected**:
   - `feat:` (conventional commits)
   - `feat()` (parentheses format)
   - `feat ` (space after feat)
   - `^feat` (feat at start of line)
3. **Feature commits**: Increment minor version (second point) and reset patch to 00
4. **Other commits**: Increment patch version (third point)
5. **Manual override**: Force specific increment types when needed

## 🚨 Safety Features

- **Branch validation**: Warns if not on main/master branch
- **Working tree check**: Ensures no uncommitted changes
- **Interactive confirmation**: Asks before proceeding with deployment
- **Error handling**: Exits on any failure to prevent partial deployments

## 📝 Example Workflow

### Daily Development
```bash
# Make your changes (various feat formats supported)
git add .
git commit -m "feat: add new endpoint"        # feat: format
git commit -m "feat(user): add authentication" # feat(scope): format
git commit -m "feat add new feature"           # feat format
git commit -m "feat() new functionality"       # feat() format

# Quick deploy (auto-detects any feat pattern → minor increment)
./quick-deploy.sh
```

### Feature Release
```bash
# Auto-deploy with feature detection
./deploy.sh

# Or force feature increment
./deploy.sh feat
```

### Bug Fixes
```bash
# Make your changes
git add .
git commit -m "fix: resolve authentication issue"

# Quick deploy (auto-detects non-feat commit → patch increment)
./quick-deploy.sh
```

### Breaking Changes
```bash
# Deploy with major version increment
./deploy.sh major
```

## 🐛 Troubleshooting

### Common Issues

1. **"Not in a git repository"**
   - Ensure you're in the project directory
   - Run `git init` if needed

2. **"Heroku remote not found"**
   - Add Heroku remote: `heroku git:remote -a YOUR_APP_NAME`

3. **"Failed to build binary"**
   - Check Go syntax: `go build`
   - Ensure all dependencies are available

4. **"Failed to deploy to Heroku"**
   - Check Heroku status: `heroku status`
   - Verify app exists: `heroku apps`

### Manual Deployment

If scripts fail, you can manually deploy:

```bash
# Build binary
go build -o main

# Commit
git add main
git commit -m "chore(release): v1.02.03"

# Tag
git tag v1.02.03

# Deploy
git push heroku master

# Push to origin
git push origin master
git push origin --tags
```

## 🔄 Updating Scripts

The scripts are designed to be maintainable. Key functions:

- `get_latest_version()` - Retrieves current version from git tags
- `increment_version()` - Handles semantic versioning logic
- `build_binary()` - Compiles Go application
- `deploy_to_heroku()` - Handles Heroku deployment
- `push_to_origin()` - Syncs with origin repository

## 📚 Additional Resources

- [Heroku Go Deployment](https://devcenter.heroku.com/articles/go-support)
- [Semantic Versioning](https://semver.org/)
- [Git Tagging](https://git-scm.com/book/en/v2/Git-Basics-Tagging)
