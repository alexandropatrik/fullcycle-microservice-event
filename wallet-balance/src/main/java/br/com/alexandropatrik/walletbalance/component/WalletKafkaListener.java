package br.com.alexandropatrik.walletbalance.component;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import br.com.alexandropatrik.walletbalance.domain.dto.AccountBalanceConsumerDto;
import br.com.alexandropatrik.walletbalance.service.AccountBalanceService;
import lombok.extern.slf4j.Slf4j;

@Component
@Slf4j
public class WalletKafkaListener {

    private final ObjectMapper objectMapper;
    private final AccountBalanceService accountBalanceService;

    public WalletKafkaListener(ObjectMapper objectMapper, AccountBalanceService accountBalanceService) {
        this.objectMapper = objectMapper;
        this.accountBalanceService = accountBalanceService;
    }

    @KafkaListener(topics = "balances")
    public void consume(@Payload String valor) {
        try {
            AccountBalanceConsumerDto dto = objectMapper.readValue(valor, AccountBalanceConsumerDto.class);
            log.info(dto.toString());
            accountBalanceService.updateBalance(dto);
        } catch (JsonMappingException e) {
            log.error(valor, e);
        } catch (JsonProcessingException e) {
            log.error(valor, e);
        }
    }

}
