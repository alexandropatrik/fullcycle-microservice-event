package br.com.alexandropatrik.walletbalance.domain.dto;

import java.io.Serializable;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
public class AccountBalanceConsumerDto implements Serializable {

    @JsonProperty("Name")
    private String name;

    @JsonProperty("Payload")
    private AccountBalanceConsumerPayloadDto payload;
    
}