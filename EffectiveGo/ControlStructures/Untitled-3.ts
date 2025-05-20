public async Task<PaymentResponse> CheckPhoneAsync(string phone)
{
    // İçeriği oluşturuyoruz
    var content = new StringContent($"{{\"phone\":\"{phone}\"}}", Encoding.UTF8, "application/json");

    // HTTP isteğini oluşturuyoruz
    var requestMessage = new HttpRequestMessage(HttpMethod.Post, $"{BaseUrl}check")
    {
        Content = content // İçeriği doğru nesneye bağlıyoruz
    };

    // İsteği gönderiyoruz
    var response = await _httpClient.SendAsync(requestMessage);

    // Yanıtı çözümlüyoruz
    return await ParseResponse<PaymentResponse>(response);
}
