CREATE DATABASE [StarWarsDB]
GO

USE [StarWarsDB]
GO

CREATE TABLE [Planets] 
(
   [Id] UNIQUEIDENTIFIER PRIMARY KEY,
   [Name] VARCHAR(100) NOT NULL,
   [Climate] VARCHAR(100) NOT NULL,
   [Terrain] VARCHAR(100) NOT NULL,
);
GO

CREATE TABLE [Films]
(
    [Id] UNIQUEIDENTIFIER PRIMARY KEY,
    [PlanetId] UNIQUEIDENTIFIER,
	[Title] VARCHAR(100) NOT NULL,
	[Director] VARCHAR(100) NOT NULL,
	[ReleaseDate] VARCHAR(100) NOT NULL,

    FOREIGN KEY (PlanetId) REFERENCES Planets(Id)
    ON DELETE CASCADE
);