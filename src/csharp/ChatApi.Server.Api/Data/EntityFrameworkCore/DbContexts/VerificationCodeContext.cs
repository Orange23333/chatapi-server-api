using Microsoft.EntityFrameworkCore;

namespace ChatApi.Server.Api.Data.EntityFrameworkCore.DbContexts
{
	public class VerificationCodeContext : DbContext
	{
		public DbSet<VerificationCodeManager.VerificationCode> verificationCodes { get; set; }
		public DbSet<VerificationCodeManager.TemporaryVerificationCode> temporaryVerificationCodes { get; set; }

		public VerificationCodeContext(DbContextOptions<VerificationCodeContext> options) : base(options) {
			;
		}
	}
}
