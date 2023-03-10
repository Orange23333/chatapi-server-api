using EleCho.Json;
using YamlDotNet;

namespace ChatApi.Server.Api.Data
{
public enum RegisterRequestStatus{
	Waiting =0,
	Accept=1,
	Decline =-1
}

	public class RegisterRequest
	{
public Guid RegisterRequestId {get;set;}

		public string UserName { get; set; }

		public string Email {get;set;}

		public (string HashType, string PasswordHash) PassWord { get; set; }

		public string Comment {get;set;}



		public RegisterRequestStatus Status {get;set;}

		public Guid? Result_UserId {get;set;}

		public Guid? HandlerAdmin_UserId {get;set;}
	}
}
