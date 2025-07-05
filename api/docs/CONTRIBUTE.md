# Contributing to Shopping List API

Thank you for your interest in contributing to the Shopping List API! This project follows [Grug's philosophy of simplicity](https://grugbrain.dev/), so we appreciate contributions that maintain this core principle.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:

   ```bash
   git clone git@github.com:your-username/shopping-list.git
   cd shopping-list/api
   ```

## Development Workflow

### Branch Strategy

- **Main Branch**: `main` - Production-ready code
- **Development Branch**: `develop` - Integration branch for new features
- **Feature Branches**: `feature/your-feature-name` - Individual feature development

### Creating a Pull Request

1. **Base your work on the `develop` branch**:

   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following the project's coding standards

3. **Test your changes**:

   ```bash
   make test
   make build
   ```

4. **Commit your changes** with clear, descriptive messages:

   ```bash
   git add .
   git commit -m "Add feature: brief description of what you added"
   ```

5. **Push to your fork**:

   ```bash
   git push origin feature/your-feature-name
   ```

6. **Create a Pull Request** against the `develop` branch (not `main`)

## Pull Request Guidelines

### Requirements

- **Target Branch**: All PRs must target the `develop` branch
- **Description**: Provide a clear description of what your PR does
- **Testing**: Include tests for new functionality
- **Documentation**: Update documentation if needed
- **Simplicity**: Follow Grug's philosophy - keep it simple

### PR Template

When creating a PR, please include:

```markdown
## Description
Brief description of the changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Refactoring

## Testing
- [ ] Tests pass locally
- [ ] New tests added (if applicable)
- [ ] Manual testing completed

## Grug's Simplicity Check
- [ ] 5-year-old could understand this change
- [ ] No unnecessary complexity added
- [ ] Follows existing patterns
- [ ] Big buttons still work
```

## Code Standards

### Go Code Style

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable names
- Keep functions small and focused
- Add comments for complex logic
- Follow existing patterns in the codebase

### Database Changes

See [DATABASE.md](docs/DATABSE.md) in this directory

- Use Goose migrations for all database changes
- Include both up and down migrations
- Test migrations thoroughly
- Document any breaking changes

### API Changes

- Maintain backward compatibility when possible
- Follow existing URL patterns
- Keep responses simple and consistent
- Document new endpoints

## Grug's Philosophy Guidelines

When contributing, remember Grug's core principles:

### ✅ Good Contributions

- Make existing features more reliable
- Fix bugs that break the "paper list" experience
- Improve performance without adding complexity
- Add essential features that families actually need
- Simplify existing code

### ❌ Avoid These (Complexity Demons)

- Adding user accounts or authentication
- Complex UI frameworks or build processes
- Features that require tutorials
- Settings or configuration pages
- Real-time notifications or fancy animations
- Integration with external services

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run specific package tests
go test ./app/db/...
```

### Manual Testing

1. Build and run the application
2. Test family creation and list management
3. Verify items can be added/removed
4. Test on mobile devices
5. Test with slow internet connection

## Documentation

### When to Update Documentation

- Adding new API endpoints
- Changing existing behavior
- Adding configuration options
- Fixing documentation errors

### Documentation Files

- `README.md` - Main project overview
- `docs/DATABASE.md` - Database schema and design
- `docs/CONTRIBUTE.md` - This file
- Code comments - For complex logic

## Getting Help

### Questions?

- Open an issue with the `question` label
- Check existing issues and PRs first
- Be specific about your environment and problem

### Reporting Bugs

- Use the bug report template
- Include steps to reproduce
- Mention your operating system and Go version
- Include relevant logs or error messages

## Review Process

1. **Automated Checks**: PRs must pass all automated tests
2. **Code Review**: At least one maintainer will review your code
3. **Testing**: Changes will be tested in a development environment
4. **Merge**: Approved PRs are merged into `develop`
5. **Release**: Periodically, `develop` is merged to `main` for releases

## Release Cycle

- **Feature Development**: Happens on `develop` branch
- **Testing**: Features are tested on `develop`
- **Release**: Stable `develop` is merged to `main`
- **Hotfixes**: Critical fixes may go directly to `main`

## Code of Conduct

- Be respectful and constructive
- Focus on the code, not the person
- Remember Grug's wisdom: simple is better
- Help maintain a welcoming community

---

Remember: **"Best shopping list app is one that feels like not using app at all"** - Ancient Grug Wisdom

Thank you for helping keep this project simple and family-friendly! 🦣
