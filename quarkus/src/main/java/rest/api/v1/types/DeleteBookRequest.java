package rest.api.v1.types;

import javax.validation.constraints.NotBlank;

public class DeleteBookRequest {

    @NotBlank
    public int id;
    
}
