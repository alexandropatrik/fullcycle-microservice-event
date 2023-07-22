package br.com.alexandropatrik.walletbalance.service;

import java.math.BigDecimal;
import java.util.UUID;

import javax.transaction.Transactional;

import org.springframework.stereotype.Service;

import br.com.alexandropatrik.walletbalance.domain.dto.AccountBalanceConsumerDto;
import br.com.alexandropatrik.walletbalance.domain.dto.AccountBalanceResourceOutputDto;
import br.com.alexandropatrik.walletbalance.domain.entity.AccountBalanceEntity;
import br.com.alexandropatrik.walletbalance.exception.ResourceNotFoundException;
import br.com.alexandropatrik.walletbalance.repository.AccountBalanceRepository;

@Service
@Transactional
public class AccountBalanceService {
    
    private final AccountBalanceRepository accountBalanceRepository;

    public AccountBalanceService (AccountBalanceRepository accountBalanceRepository) {
        this.accountBalanceRepository = accountBalanceRepository;
    }

    public void updateBalance(AccountBalanceConsumerDto accountBalanceConsumerDto) {
        updateBalance(accountBalanceConsumerDto.getPayload().getAccountIdFrom(), accountBalanceConsumerDto.getPayload().getBalanceAccountIdFrom());
        updateBalance(accountBalanceConsumerDto.getPayload().getAccountIdTo(), accountBalanceConsumerDto.getPayload().getBalanceAccountIdTo());
    }

    public void updateBalance(UUID accountBalanceId, BigDecimal balance) {
        AccountBalanceEntity accountBalanceEntity = accountBalanceRepository
            .findById(accountBalanceId)
            .orElse(AccountBalanceEntity.builder().accountId(accountBalanceId).build());
        accountBalanceEntity.setBalance(balance);
        accountBalanceRepository.save(accountBalanceEntity);
    }

    public AccountBalanceResourceOutputDto findBalanceById(UUID accountId) throws ResourceNotFoundException {
        AccountBalanceEntity accountBalanceEntity = accountBalanceRepository
            .findById(accountId)
            .orElseThrow(ResourceNotFoundException::new);

        return AccountBalanceResourceOutputDto.builder()
            .accountId(accountId)
            .balance(accountBalanceEntity.getBalance())
            .build();
    }

}
