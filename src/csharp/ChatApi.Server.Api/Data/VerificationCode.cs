namespace ChatApi.Server.Api.Data
{
	public static class VerificationCode
	{
		private static SortedDictionary<Guid, string> verificationCodes = new SortedDictionary<Guid, string>();
		private static SortedDictionary<Guid, DateTime> temporaryVerificationCodeIds = new SortedDictionary<Guid, DateTime>();

		public static int TemporaryVerificationCodeIdLimit { get; set; } = 1024;
		public static int VerificationCodeLimit { get; set; } = 1024;

		public static void CleanOverdueTemporaryVerificationCode(TimeSpan timeOut)
		{
			DateTime now = DateTime.UtcNow;
			int i;

			lock(temporaryVerificationCodeIds)
			{
				for (i = 0; i < temporaryVerificationCodeIds.Count;)
				{
					var item = temporaryVerificationCodeIds.ElementAt(i);
					if(now - item.Value >= timeOut)
					{
						temporaryVerificationCodeIds.Remove(item.Key);
					}
					else
					{
						i++;
					}
				}
			}
		}

		public static Guid? GetOne_TemporaryVerificationCodeId()
		{
			if(temporaryVerificationCodeIds.Count < TemporaryVerificationCodeIdLimit)
			{
				Guid newId = Guid.NewGuid();

				// Ignore duplicated GUID.
				temporaryVerificationCodeIds.Add(newId, DateTime.UtcNow);

				return newId;
			}

			return null;
		}

		public static Guid? GetOne(Guid tempVCId)
		{
			if(verificationCodes.Count < VerificationCodeLimit)
			{

			}
		}
	}
}
