using EleCho.Json;
using YamlDotNet;

namespace ChatApi.Server.Api.Data
{
	public class UserProfile
	{
		public Guid UserId { get; set; }

		public string UserName { get; set; }

		public (string HashType, string PasswordHash) PassWord { get; set; }

		public 
	}
}
