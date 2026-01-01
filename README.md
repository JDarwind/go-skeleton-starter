Go skeleton Repository

# Important
This project is maintained by a third party and is not affiliated in any way with the Go authors or maintainers.

This repository contains opinionated code intended to speed up the bootstrapping of your application. Nevertheless, this repository is not a complete or standalone application and must not be considered production-ready by design.

Before using any part of this code in production, make sure you fully understand which components are required and whether they fit your project’s requirements.

# Using this repository as template

To use this repository, you need to have the gonew tool installed.

Once installed, you can run the following command:

```bash
gonew github.com/JDarwind/go-skeleton-starter <your-namespace> <path-to-your-project>
```

## Example 
Assume your namespace is:
```bash
github.com/your_user/your_awesome_project
```
You can run:

```bash
gonew github.com/JDarwind/go-skeleton-starter github.com/your_user/your_awesome_project
```
This will create a folder named `your_awesome_project` in the directory where you ran the command.

## Example (Current Directory)
If you want to create the project in the current directory instead of a subfolder, you can run:

```bash
gonew github.com/JDarwind/go-skeleton-starter github.com/your_user/your_awesome_project .
```

# Installing GONEW
To install gonew, follow the installation instructions in the official documentation available at this link:

https://pkg.go.dev/golang.org/x/tools/cmd/gonew

# NOTES
- The official documentation link may change over time, as this project is not maintained or affiliated in any way with the Go authors or maintainers.

- Support for any tool such as `gonew` will not be provided nor guaranteed, since there is no affiliation between this project, or it's author and the gonew authors.

- `gonew` is used in this README only as an example. Other alternatives may exist or be introduced in the future.