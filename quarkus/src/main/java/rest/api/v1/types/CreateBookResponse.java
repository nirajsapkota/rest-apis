package rest.api.v1.types;

public class CreateBookResponse {
    private String message;
    
    public CreateBookResponse(String message) {
        this.message = message;
    }
    
    public String getMessage() {
        return message;
    }
    
    public void setMessage(String message) {
        this.message = message;
    }
}
