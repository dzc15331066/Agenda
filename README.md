# Agenda

# Introduction
``Agenda`` is a simple meeting manegement system developed with`` Golang`` and it's a powerful modern CLI application as well. As a wonderful simple project to practise coding in Golang, Agenda offers many functions including regsitering an user, logging into system, logout from systen, creating a meeting with meeting's title, meeting's participators, and meeting's starting time and ending time, besides, you can add or delete participartors in some meeting. Of course, Agenda supports querying some user or meetings, and other funstions wait for your exploration.

# Overview
Agenda is`` CLI ``application developed with the library called ``Cobra``, which provides a simple interface to create powerful modern CLI interfaces similar to git & fo tools. Thanks to Cobra, we can finish this project within a short time, and spent less time in parsing commans but concentrate our attention on the logical bussiness layer and data layer. parts of help of Cobra offers are listes as follow:

### Cammands
Under the help of Cobra, we can create some commands easily, for example, in Agenda, these Commands are created.


### Flags

Under the help of Cobra, wen can easily register and parse the commands and create subCommands, for example, in our application, some Flags are parsed as follow:

register command of login:
```
	loginCmd.Flags().StringP("user", "u", "", "use -user [username] or -u [username]")
	loginCmd.Flags().StringP("password", "p", "", "user -password [password] or -p [password]")
```
parse it

```
	username, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
```
### MVC and oriented object design mode

Agenda use three layer (view, logical controller, data) to build up the project and use oriented object designing mode to achieve entities like AgendaService, User, Meeeting, which looks as if those entities are real objects compared with obeject in Java or C++.



# Design
### Achievements of course requires
* Use ``Json`` file to read or write the entities , User and  Meeting.
* Support service of ``log`` ,records and traces the operations of user and some important output. 
* Use ``.travis.yml ``file in project
* Support ``automatical testing``, help you tesing the apllication quickly.
* Use ``Flags`` to parse commands
* Use ``Time`` to parse time

### Designing of Agenda
  we use Architecture of`` MVC``, achieves the separetion among ``view`` layer, ``logical controll`` layer and ``data layer``. 
  
    Agenda/cmd/
  
    addPart.go
    clear.go	
    cm.go	
    delPart.go	
    delUser.go	
    dm.go	
    em.go	
    login.go	
    logout.go
    qm.go	
    query.go	
    register.go	
    root.go
    
 Above are regared as the view layer, which focus on receiving and handling user commands. for each commands, it calls related methods
of AgendaService to interate with data layer. for example, the command of ``login``,which just need to call APIs from AgendaService.
```
var loginCmd = &cobra.Command{
  Use:   "login",
  Short: "A brief description of your command",
  Long: `A longer description that spans multiple lines and likely contains examples
 and usage of using your command. For example:
 Cobra is a CLI library for Go that empowers applications.
 This application is a tool to generate the needed files
 to quickly create a Cobra application.`,
 
	Run: func(cmd *cobra.Command, args []string) {
		log.Infoln("Login:")
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		err := as.UserLogin(username, password)
		message(err, "[success]: login successfully!")
	},
}
```
we regard agendaservice and storage as entities too, and put them together with User and Meeting under the directory of entity.

```
    Agenda/entity/
    
    agendaService.go	
    meeting.go	
    storage.go	
    user.go
    
 ```
 The AgendaService.go contains all the APIs for commands to interate with data layer, for example, ``login`` for user, which just provides API for commands to operate data but never do any operation directly with data.
 
 ```
  func (as *AgendaService) UserLogin(username string, password string) error {
    if username == "" || password == "" {
     return nullAgumentError
    }
    if err := as.AgendaStorage.readUsers(); err != nil {
     return err
    }
    res := as.AgendaStorage.QueryUser(func(user User) bool {
     return username == user.Name && password == user.Password
    })

    if len(res) > 0 {
     return as.AgendaStorage.setCurUser(res[0])

    }
    return errors.New("[error]: Invalid username or password")
  }
 ```
 Ok, the operation will effects once the data in file is modified. we use setCurUser method to record the current user once the user login and write the record out to curUser.txt file as the state of login, which provides help for next command.
 
 ```
 func (s *storage) setCurUser(user User) error {
   s.CurUser = user
   return s.writeCurUser()
 }
 
 func writeToFile(datalist interface{}, filename string) error {
   data, err := json.Marshal(datalist)
   if err != nil {
    return err
   }
   return ioutil.WriteFile(filename, data, 0666)
 }

 ```
 
# Testing

In this module, we've finished testing work for each APIs of AgendaService and now we can conclude that everything works well. Of course, we use the golang pakeage of ``testing ``, which helps a lot in ``unit test``. some of our testing results are showed as follow:

here is the testing results


# Installing
Using Agenda is easy.First, use ``git clone`` to install the project

    git clone https://github.com/dzc15331066/Agenda
    
# Getting Started
 While you are welcome to provide your own organization, typically our Agenda apllication shows the following organizaion structure.
 ```
 Agenda
 --cmd
 	--addPart.go
    	--clear.go	
    	--cm.go	
    	--delPart.go	
    	--delUser.go	
    	--dm.go	
    	--em.go	
    	--login.go	
    	--logout.go
    	--qm.go	
    	--query.go	
    	--register.go	
    	--root.go
 --entity
 	--agendaService.go	
   	--meeting.go	
    	--storage.go	
    	--user.go
 --test
 	--agenda_test.go
	--test.sh
```

To start ``Agenda``, you should enter the directory ``Agenda/``, and run the command ``go build``, and then you can use Agenda commands as this format ``Agenda [Command] [subCommand]...``, more details can be get once you enter ``./agenda -h``, enjoy yourself!



[git协同开发参考](https://github.com/livoras/blog/issues/7)
