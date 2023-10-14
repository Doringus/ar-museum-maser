package com.monkeprojects.armaster.domain;

import org.bson.types.Binary;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.Date;

@Document(collection = "Exhibitions")
public class Exhibition {
    @Id
    private String id;
    private String name;

    public Exhibition(String name) {
        this.name = name;
    }

    public String getId() {
        return id;
    }


    @Override
    public boolean equals(Object object) {
        if (this == object) {
            return true;
        }
        if (object == null || getClass() != object.getClass()) {
            return false;
        }
        Exhibition exhibition = (Exhibition) object;
        return  java.util.Objects.equals(name, exhibition.name)
                && java.util.Objects.equals(id, exhibition.id);
    }

    @Override
    public int hashCode() {
        return java.util.Objects.hash(name, id);
    }

    @Override
    public String toString() {
        return "Exhibition {"
                + "name='" + name + '\''
                + ", id='" + id + '\''
                + '}';
    }

}
