using Grpc.Core;

namespace BookshopServer.Services;

public class InventoryService : Inventory.InventoryBase
{
    public override Task<GetBookDetailByIdResponse> GetBookDetailById(GetBookDetailByIdRequest request, ServerCallContext context)
    {
        return base.GetBookDetailById(request, context);
    }

    public override Task<GetBookListResponse> GetBookList(GetBookListRequest request, ServerCallContext context)
    {
        var response = new GetBookListResponse();

        response.Books.Add(new Book
        {
            Title = "Tes"
        });

        return Task.FromResult(response);
    }
}