# 📚 Golang AWS OpenSearch Index Management

🚀 Golang snippets demonstrating how to connect to **AWS OpenSearch** (formerly known as Amazon Elasticsearch) with IAM-based role authentication.

### 📄 Files
- **`openserverless_serverless_conn_create_index.go`**: Create and delete indexes with custom settings.
- **`openserverless_serverless_conn_get_index.go`**: Retrieve index metadata.

## ✅ Prerequisites

Before you begin, make sure you have:

1. 🛡️ **An AWS Account**: With IAM roles and permissions configured to access AWS OpenSearch.
2. 🛠️ **Golang Installed**: Ensure Golang is set up on your machine.

## ⚙️ Setup

1. **Set Environment Variables**:

   Configure the following environment variables to enable the code to interact with your OpenSearch cluster:

   - `AWS_OPENSEARCH_ENDPOINT`: 🌐 Your AWS OpenSearch cluster's endpoint URL.
   - `AWS_REGION`: 🌎 The AWS region where your OpenSearch cluster is hosted.

---

⚠️ **Important**: Index creation and deletion are powerful operations that directly impact your AWS OpenSearch data. Proceed with caution to avoid accidental data loss.
