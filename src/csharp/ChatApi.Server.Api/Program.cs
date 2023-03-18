using ChatApi.Server.Api.Data.EntityFrameworkCore.DbContexts;
using Microsoft.AspNetCore.Hosting;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Configuration;
using System.Linq;
using System.Threading.Tasks;

namespace ChatApi.Server.Api
{
	public class Program
	{
		public static void Main(string[] args)
		{
			var builder = WebApplication.CreateBuilder(args);

			// Add services to the container.
			builder.Services.AddAuthorization();

			// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
			builder.Services.AddEndpointsApiExplorer();
			builder.Services.AddSwaggerGen();

			builder.Services.AddDbContext<VerificationCodeContext>(options =>
			{
				// appsettings.json:
				//  "ConnectionStrings": {
				//   "DatabaseConnection": "server=localhost;database=chatapi;uid=chatapi;pwd=chatapi123;"
				//  }
				//options.UseMySql(Configuration.GetConnectionStrings("DatabaseConnecton"));
				options.UseMySQL(connectionString: "server=localhost;database=chatapi;uid=chatapi;pwd=chatapi123;");
			});
			builder.Services.AddControllersWithViews();

			var app = builder.Build();

			// Configure the HTTP request pipeline.
			if (app.Environment.IsDevelopment())
			{
				app.UseSwagger();
				app.UseSwaggerUI();
			}

			app.UseHttpsRedirection();

			app.UseAuthorization();

			app.Run();
		}
	}
}