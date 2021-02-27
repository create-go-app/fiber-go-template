# /platform

This folder contains all the platform-level logic that will build up the actual product, like setting up the database, Redis instance, or maybe a router.

All of the packages inside this directory are independent of each other.

Packages inside this folder also don't care what's inside `/app` or `/pkg`.
