# Agenda

# Introduction
``Agenda`` is a simple meeting manegement system developed with`` Golang`` and it's a powerful modern CLI application as well. As a wonderful simple project to practise coding in Golang, Agenda offers many functions including regsitering an user, logging into system, logout from systen, creating a meeting with meeting's title, meeting's participators, and meeting's starting time and ending time, besides, you can add or delete participartors in some meeting. Of course, Agenda supports querying some user or meetings, and other funstions wait for your exploration.

# Overview
Agenda is`` CLI ``application developed with the library called ``Cobra``, which provides a simple interface to create powerful modern CLI interfaces similar to git & fo tools. Thanks to Cobra, we can finish this project within a short time, and spent less time in parsing commans but concentrate our attention on the logical bussiness layer and data layer. parts of help of Cobra offers are listes as follow:

### Cammands
Under the help of Cobra, we can create some commands easily, for example, in Agenda, these Commands are created.


### Flags

Under the help of Cobra, wen can easily parse the commands and create subCommands, for example, in our application, some Flags are parsed as follow:

# Design
### Achievements of course requires
* Use ``Json`` file to read or write the entities , User and  Meeting.
* Support service of ``log`` ,records and traces the operations of user and some important output. 
* Have ``.travis.yml ``file in project
* Support ``automatical testing``, help you tesing the apllication quickly.
* Use ``Flags`` to parse commands

### designing of Agenda
  we use Architecture of MVC, achieves the separetion among view layer, logical controll layer and data layer. 

# Testing

# Installing

# Getting Started



[git协同开发参考](https://github.com/livoras/blog/issues/7)
