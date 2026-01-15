package com.example.chat.security;

import java.io.IOException;

import org.springframework.stereotype.Component;

import jakarta.servlet.Filter;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.ServletRequest;
import jakarta.servlet.ServletResponse;
import jakarta.servlet.http.HttpServletRequest;

@Component
public class JWTAuthFilter implements Filter {
    private final JWTValidator validator;

    public JWTAuthFilter(JWTValidator validator) {
        this.validator = validator;
    }

    @Override
    public void doFilter(
            ServletRequest req,
            ServletResponse res,
            FilterChain chain
    ) throws IOException, ServletException {
        HttpServletRequest request = (HttpServletRequest) req;
        String header = request.getHeader("Authorization");
        
        if (header != null && header.startsWith("Bearer ")) {
            String token = header.substring(7).trim();
            try {
                String userId = validator.validateAndGetUserId(token);
                request.setAttribute("userId", userId);
                request.setAttribute("token", token);
            } catch (Exception e) {
                throw new ServletException("Invalid token " + e.getMessage());
            }
        }
        chain.doFilter(req, res);
    }
}
