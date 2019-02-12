# FS Abstraction Layer

|     |     |
| -   |   - |
| **Project status** | Proposal |
| **API Status** | Not frozen |

## Reasons

In development you need to manage file uploads in different environments. Different environments means different disks, raids or even different storage providers.

But every storage provider provides same features and ideas: the same "read file", "open dir", "created at", "modified at", "create file", etc.
