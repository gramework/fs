# FS Abstraction Layer

|     |     |
| -   |   - |
| **Project status** | Proposal |
| **API Status** | Not frozen |

## Reasons

In development you need to manage file uploads in different environments. Different environments means different disks, raids or even different storage providers.

But every storage provider provides same features and ideas: the same "read file", "open dir", "created at", "modified at", "create file", etc.

I didn't find a solution made in a way I need it, so I decided to build my own.

**Main idea of FS Abstraction Layer** is to provide simple library that will allow to plug-in complex storage providers like Akamai NetStorage or AWS S3 without need to change any source code except configuration files.
