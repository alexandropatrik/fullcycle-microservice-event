package br.com.alexandropatrik.walletbalance.repository;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import br.com.alexandropatrik.walletbalance.domain.entity.AccountBalanceEntity;

public interface AccountBalanceRepository extends JpaRepository<AccountBalanceEntity, UUID> {
    
}
