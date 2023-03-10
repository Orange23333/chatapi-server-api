using EleCho.Json;
using YamlDotNet;

namespace ChatApi.Server.Api.Data
{
	public class VerfiyCodeRequest
	{
		public Guid UserId { get; set; }

		public List<Guid> RequestIds {get;set;}
	}
}
