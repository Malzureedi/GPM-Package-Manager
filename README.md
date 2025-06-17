GPM, The Unofficial Git Package Manager.

GPM works by Cloning Git (including GitHub or Gitlab repository) repositories - it mainly uses HTTPS to clone.

Why use GPM?

- Fast; we use Golang which is a Compiled language which mean's its closer to hardware meaning more peformance.
- Git; it uses Git by Linus Torvalds to Clone repositories.
- Simple; This heavily simplifies the Git library for Package cloning and Installing.

Just be mindful of what Packages your installing and what Repository it is.
Reasons:
- Git Repo's do not have much Security, what I mean is malware and Malicious Program detection.
- Some Git repositories are NOT the real one, make sure your only using the real repository.
- Some Git Repository's may not be compatible with GPM Autobuild, meaning some will need manual building/configs.

GPM Autobuild:
This software is apart of the gpm-any package (any is any architecture.) that will compile the program.
Some programs may be not compatible with your architecture, Gpm autobuild cannot change architecture code.

It works by checking the Package for a "gpmbuild.json" that contains instructions and dependency lists for GPM.

GPM is primarily a Linux Package Manager.

How to Build:
Just type in make, as long as you have all the Depenedency's installed it should work!

Dependency's:
-- Git
-- Python3
-- Go

# Usage

To install a Package you can use;
gpm -i (repository git link here)
You can also find Git Repository's by using:
gpm -f (git repository link here)
To remove Package files (uses rm) then you can do this;
gpm -r (filename)

