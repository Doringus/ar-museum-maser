package com.monkeprojects.armaster.domain;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.Date;

@Document(collection = "file")
public class File {
    @Id
    private String id;
    private String name;
    private Date uploadDate;
    private byte[] content;
    private long size;
    protected File() {
    }

    public File(String name, long size, byte[] content) {
        this.name = name;
        this.uploadDate = new Date();
        this.content = content;
        this.size = size;
    }

    public byte[] getContent() {
        return content;
    }

    public String getId() {
        return id;
    }
    public String getName() { return name;}
    public long getSize() { return size;}


    @Override
    public boolean equals(Object object) {
        if (this == object) {
            return true;
        }
        if (object == null || getClass() != object.getClass()) {
            return false;
        }
        File fileInfo = (File) object;
        return  java.util.Objects.equals(name, fileInfo.name)
                && java.util.Objects.equals(uploadDate, fileInfo.uploadDate)
                && java.util.Objects.equals(id, fileInfo.id)
                && java.util.Objects.equals(size, fileInfo.size);
    }

    @Override
    public int hashCode() {
        return java.util.Objects.hash(name, uploadDate, id, size);
    }

    @Override
    public String toString() {
        return "File{"
                + "name='" + name + '\''
                + ", uploadDate=" + uploadDate
                + ", size =" + size
                + ", id='" + id + '\''
                + '}';
    }

}
