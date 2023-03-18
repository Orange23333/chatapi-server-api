using ChatApi.Server.Api.Data;
using Microsoft.AspNetCore.Mvc;
using static ChatApi.Server.Api.Data.VerificationCodeManager;

namespace ChatApi.Server.Api.Controllers
{
	[ApiController]
	[Route("verify-code/[controller]")]
	public class VerificationCodeController
	{
		[HttpGet]
		public static Guid? GetOne()
		{
			return VerificationCodeManager.CreateTemporaryVerificationCode();
		}

		[HttpGet]
		public static Guid? GetOne(Guid tempVCId)
		{
			return VerificationCodeManager.CreateVerificationCode(tempVCId);
		}
	}
}
