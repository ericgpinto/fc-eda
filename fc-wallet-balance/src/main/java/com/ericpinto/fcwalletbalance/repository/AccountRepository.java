package com.ericpinto.fcwalletbalance.repository;

import com.ericpinto.fcwalletbalance.entity.AccountEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AccountRepository extends JpaRepository<AccountEntity, String> {
}
