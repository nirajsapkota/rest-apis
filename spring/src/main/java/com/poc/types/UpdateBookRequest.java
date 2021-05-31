package com.poc.types;

import javax.validation.constraints.NotNull;

public class UpdateBookRequest {
    
    @NotNull
    public int id;
    
    public String title;
    
    public String author;
    
}
