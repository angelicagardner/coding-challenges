# Coding Challenges

This repository saves my solutions to the following different coding challenges:

- [Cracking The Coding Interview (CTCI)](https://github.com/careercup/CtCI-6th-Edition)
- [Leetcode](https://leetcode.com/)
- [Project Euler](https://projecteuler.net/archives)

The solutions are written in Python.


## Format

1. Problem/Challenge is created as an Issue. The issue will have relevant associated Labels.
2. Solution to a problem/challenge (issue) is submitted as a Pull Request. The solution has an associated file with tests that will be used to verify the solution.
3. The solution file will be named to reflect problem/challenge (issue) name, e.g. **problem_001.py**, and and associated test file will have the same name and end with **_test** suffix, e.g. **problem_001_test.py**.
4. The solution file should be structured according to the template below:

### Solution Template

```python
"""
[Issue title]: [link]

[Short solution explanation]
"""

import module1, module2

# Helper functions are optional
def helperFunction(arg1: [type hint], arg2: [type hint], ...) -> [Return type hint]:
    ...

    return


def solution(arg1: [type hint], arg2: [type hint], ...) -> [Return type hint]:
    ...

    return answer


if __name__ == "__main__":
    print(f"{solution() = }")
```


## Set up

The repository is organization into folders for the different projects. Each created issue will belong to one of the challenges. Once a coding solution is verified, it is merged and placed in the relevant folder - together with the associated test file. The solutions are checked/verified by automated testing in a GitHub Actions workflow. The efficiency of the code is also checked. The workflow is triggered when a pull request is opened or updated.


### Prerequisities

For the setup to work, the repository needs to have a **Actions secret** created in **Settings**. The secret will be a pasted personal access token (PAT) generated from GitHub Profile **Developer settings**.