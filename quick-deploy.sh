#!/bin/bash

# Quick deploy script for Go API on Heroku
# Simple version that just increments patch and deploys

set -e

echo "🚀 Quick Deploy to Heroku"

# Get latest version and determine increment type
current_version=$(git tag --sort=-version:refname | head -1)
if [[ -z "$current_version" ]]; then
    current_version="v1.00.00"
fi

# Check if there are any feat commits since last tag
last_tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
if [[ -n "$last_tag" ]]; then
    # Look for various feat commit patterns: feat:, feat(), feat, etc.
    feat_commits=$(git log --oneline "$last_tag"..HEAD | grep -E -c "(feat:|feat\(|feat\s|^feat)" 2>/dev/null || echo "0")
    if [[ "$feat_commits" -gt 0 ]] 2>/dev/null; then
        echo "🔍 Detected $feat_commits feature commit(s) - incrementing minor version"
        # Feature commits increment minor version and reset patch to 00
        IFS='.' read -ra version_parts <<< "${current_version#v}"
        major=${version_parts[0]}
        minor=${version_parts[1]}
        new_minor=$((minor + 1))
        new_version="v${major}.$(printf "%02d" $new_minor).00"
    else
        echo "🔍 No feature commits detected - incrementing patch version"
        # Other commits increment patch version
        IFS='.' read -ra version_parts <<< "${current_version#v}"
        major=${version_parts[0]}
        minor=${version_parts[1]}
        patch=${version_parts[2]}
        new_patch=$((patch + 1))
        new_version="v${major}.$(printf "%02d" $minor).$(printf "%02d" $new_patch)"
    fi
else
    # No previous tags, start with patch increment
    IFS='.' read -ra version_parts <<< "${current_version#v}"
    major=${version_parts[0]}
    minor=${version_parts[1]}
    patch=${version_parts[2]}
    new_patch=$((patch + 1))
    new_version="v${major}.$(printf "%02d" $minor).$(printf "%02d" $new_patch)"
fi

echo "📦 Current version: $current_version"
echo "🆕 New version: $new_version"

# Build binary
echo "🔨 Building Go binary..."
go build -o main

# Commit and tag
echo "💾 Committing new version..."
git add main
git commit -m "chore(release): $new_version"
git tag "$new_version"

# Deploy to Heroku
echo "🚀 Deploying to Heroku..."
git push heroku master

# Push to origin
echo "📤 Pushing to origin..."
git push origin master
git push origin --tags

echo "✅ Deployment completed! Version $new_version is live on Heroku"
