# Verifying an uploaded file is the allowed type

1. When a file is uploaded, the request headers will have a Content-Type field with either a specific content type, such as image/png, or a general value of application/octet.

```
file, header, err := r.FormFile("file")
contentType := header.Header["Content-Type"][0]
```


2. A file extension is associated with a MIME type and can provide insight into the type of file being uploaded.


3. Parsing the file and detecting the content type based on the contents of the file.

```
file, header, err := r.FormFile("file)
buffer := make([]byte, 512)
_, err = file.Read(buffer)
filetype := http.DetectContentType(buffer)
```
The buffer is only 512 bytes because `DetectContentType` looks at only up to the first 512 bytes when determining the type. When it isn't able to detect a specific type, application/octet-stream is returned.

