using System;
using System.Collections;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using Microsoft.AspNetCore.Mvc;
using SkiaSharp;

namespace ChatApi.Server.Api.Data
{
	public static class VerificationCodeManager
	{
		#region ===Classes===
		public enum VerificationCodeIdType
		{
			Other = -1,
			None = 0,
			VerificationCode = 1,
			TemporaryVerificationCode = 2
		}

		public class VerificationCode
		{
			[Key]
			public Guid Id { get; }

			public DateTime CreateTime { get; }

			public (VerificationCodeIdType, Guid)? BoundVerificationCode { get; private set; }

#warning May use HTMLCode and MetaSource to record them later.
			public SKImage? VerificationCodeImage { get; set; }

			public void BindVerificationCode(VerificationCodeIdType verificationCodeIdType, Guid guid)
			{
				if(BoundVerificationCode == null)
				{
					BoundVerificationCode = new (verificationCodeIdType, guid);
				}
				else
				{
					throw new InvalidOperationException("Has been bound a verfication code id.");
				}
			}

			public VerificationCode(Guid id) : this(id, DateTime.UtcNow, null) {; }
			public VerificationCode(Guid id, DateTime createTime, SKImage? verificationCodeImage)
			{
				this.Id = id;
				this.CreateTime = createTime;
				this.VerificationCodeImage = verificationCodeImage;
			}
		}

		public class TemporaryVerificationCode
		{
			[Key]
			public Guid Id { get; }

			public DateTime CreateTime { get; }

			public Guid? BoundVerificationCodeId { get; }

			public void BindVerificationCode(Guid verificationCodeId)
			{
				if(BoundVerificationCodeId == null)
				{
					VerificationCode verificationCode = verificationCodes[verificationCodeId];
					if(verificationCode.)
				}
				else
				{
					throw new InvalidOperationException("Has bound.");
				}
			}

			public TemporaryVerificationCode(Guid id) : this(id, DateTime.UtcNow, null) {; }
			public TemporaryVerificationCode(Guid id, DateTime createTime, Guid? boundVerificationCode)
			{
				this.Id = id;
				this.CreateTime = createTime;
				this.BoundVerificationCodeId = boundVerificationCode;
			}
		}
		#endregion

		#region ===Properties===
		public static int TemporaryVerificationCodeIdLimit { get; set; } = 1024;
		public static int VerificationCodeLimit { get; set; } = 1024;
		#endregion

		#region ===TemporaryVerificationCode===
		private static SortedDictionary<Guid, TemporaryVerificationCode> temporaryVerificationCodeIds = new SortedDictionary<Guid, TemporaryVerificationCode>();

		public static void CleanOverdueTemporaryVerificationCode(TimeSpan timeOut)
		{
			DateTime now = DateTime.UtcNow;
			int i;

			lock(temporaryVerificationCodeIds)
			{
				for (i = 0; i < temporaryVerificationCodeIds.Count;)
				{
					var item = temporaryVerificationCodeIds.ElementAt(i);
					if(now - item.Value.CreateTime >= timeOut)
					{
						if (item.Value.BoundVerificationCodeId != null)
						{
							verificationCodes[item.Value.BoundVerificationCodeId.Value]
						}

						temporaryVerificationCodeIds.Remove(item.Key);
					}
					else
					{
						i++;
					}
				}
			}
		}

		public static Guid? CreateTemporaryVerificationCode()
		{
			if(temporaryVerificationCodeIds.Count < TemporaryVerificationCodeIdLimit)
			{
				Guid newId = Guid.NewGuid();

				// Ignore duplicated GUID.
				temporaryVerificationCodeIds.Add(newId, new TemporaryVerificationCode(newId));

				return newId;
			}

			return null;
		}
		#endregion

		#region ===VerificationCode===
		private static SortedDictionary<Guid, VerificationCode> verificationCodes = new SortedDictionary<Guid, VerificationCode>();

		public static void DestoryVerificationCode()
		{

		}

		public static Guid? CreateVerificationCode(Guid tempVCId)
		{
			if (verificationCodes.Count < VerificationCodeLimit)
			{
				Guid newId = Guid.NewGuid();

				verificationCodes.Add(newId, new VerificationCode(newId));

				

				return newId;
			}
			return null;
		}
		#endregion
	}
}
