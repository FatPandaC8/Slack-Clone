package com.example.chat.web;

import java.util.List;

public record CreateChannelHttpRequest(
    String channelId,
    List<String> memberIds
) {
    
}
