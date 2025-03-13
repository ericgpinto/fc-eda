package com.ericpinto.fcwalletbalance.client;


import com.ericpinto.fcwalletbalance.dto.TransactionMessage;
import com.ericpinto.fcwalletbalance.service.TransactionService;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class TransactionConsumer {

    private static final Logger log = LoggerFactory.getLogger(TransactionConsumer.class);
    private final ObjectMapper objectMapper;
    private final TransactionService transactionService;

    public TransactionConsumer(ObjectMapper objectMapper, TransactionService transactionService) {
        this.objectMapper = objectMapper;
        this.transactionService = transactionService;
    }

    @KafkaListener(topics = "transactions", groupId = "wallet")
    @Transactional
    public void consumeTransaction(ConsumerRecord<String, String> record) {
        try {
            log.info("Received transaction event: {}", record.value());

            TransactionMessage message = objectMapper.readValue(record.value(), TransactionMessage.class);

            transactionService.processTransaction(
                    message.getPayload().getAccountIdFrom(),
                    message.getPayload().getAccountIdTo(),
                    message.getPayload().getAmount()
            );

            log.info("Transaction successfully processed.");

        } catch (Exception e){
            log.error("Error processing transaction", e);
        }
    }
}
