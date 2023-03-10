using EleCho.Json;
using YamlDotNet;

namespace ChatApi.Server.Api.Data
{
	public class UserRequests
	{
		public Guid UserId { get; set; }

		public List<Guid> RequestIds {get;set;}
	}
}
