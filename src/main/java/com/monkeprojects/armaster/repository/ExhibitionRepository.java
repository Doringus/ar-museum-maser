package com.monkeprojects.armaster.repository;

import com.monkeprojects.armaster.domain.Exhibition;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface ExhibitionRepository extends MongoRepository<Exhibition, String> {
}