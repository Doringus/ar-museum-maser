package com.monkeprojects.armaster.repository;

import com.monkeprojects.armaster.domain.File;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface FileRepository extends MongoRepository<File, String> {
}