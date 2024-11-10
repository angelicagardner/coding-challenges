# Coding Challenges

This repository saves my solutions to the following different coding challenges:

- [Cracking The Coding Interview (CTCI)](https://github.com/careercup/CtCI-6th-Edition)
- [Interviewing.io](https://interviewing.io/)
- [Leetcode](https://leetcode.com/)
- [Project Euler](https://projecteuler.net/archives)

The solutions are written in Golang.

## Format

1. Problem from one coding challenge is created as an Issue. The issue has relevant Labels and a description of the problem.

2. Solution to a Issue (coding problem) is submitted as a Pull Request. The solution also has an associated _test file with test cases to verify the solution.

3. The solution file will be named to reflect problem name, e.g. **problem_001.go**, and associated test file will have the same name and end with **_test** suffix, e.g. **problem_001_test.go**.

4. The solution file should contain a brute force solution and an improved solution. Write the time and space complexities as comments about the function. The test cases should test both solutions.

### Solution Template

```go
/*
[Issue title]
*/

package challenge

// Time complexity:     O(n^2)
// Space complexity:    O(1)
func BruteForceSolution() {
    // Implementation here
    return answer
}

// Time complexity:     O(n)
// Space complexity:    O(n)
func ImprovedSolution() {
    // Implementation here
    return answer
}
```

## Set up

The repository is organised into folders for the different challenges (projects) and each Issue created will belong to one of the challenges, i.e. One Issue = One Problem.

Once a coding solution is verified, it is merged and placed in the relevant folder - together with the associated test file. The test cases are run in a GitHub Action and the workflow is triggered when a pull request is opened or updated.

### Use Conventional Commits

This repo automatically lints commit messages with [commitlint](https://commitlint.js.org/) when new pull requests are created and merged to the main branch. To make sure the commit messages are correct, this can be set up locally.

1. Clone the repository

```bash
git clone https://github.com/angelicagardner/coding-challenges.git
cd coding-challenges
```

2. Make sure you have Node.js and npm installed

In our GitHub Action we're using Node.js v20.14.0.

3. Husky is used to manage Git hooks

Make sure husky is installed and set it up:

```bash
npx husky init
```

4. Configure Husky Hooks

Remove the "npm test" line from .husky/pre-commit file, and add the following line to the file .husky/commit-msg:

> npm run commitlint ${1}

5. Verify

You can check the last commit and this will return an error if invalid or a positive output if valid:

```bash
npx commitlint --from HEAD~1 --to HEAD --verbose
```

If everything is set up correctly you will be able to "git commit" with a commit message according to the Conventional Commits specification. Read more here: [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
