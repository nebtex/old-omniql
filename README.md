## omniql

omniql a new way to create schemas, for large project, that can be used in any devices, internally or externally, 
collaborate and create api and export them to the world


Keys aspect:
 - we use crqs (Command and Query Responsibility Segregation) if you don't know what is this please read https://martinfowler.com/bliki/CQRS.html
 - all the communication should be over nats or kafka or similar software (centrifuge for the frontend)
 - it should use the open api 3 spec https://github.com/OAI/OpenAPI-Specification/blob/3.0.0-rc0/versions/3.0.md, like initial guide when is possible
 - also we should use the google api design spec https://cloud.google.com/apis/design/

what omniql should offer:
 - creation of schemas in a nice interface
 - auto-generation of code for diferent programming languages

plan:
 - do an initial version, working
 - build omniql on omniql
 
clients:
 - large organization that has many team of different backgrounds
 - microservice infrastructures
 - iot

Why is this project needed?

is still difficult to create good schema with the current tech, if we have a good tool for create schemas, this will decrease the work needed to get things done, 
in organization that worry about infrastructure evolution. more easy work, more people happy


similar software:

- apache thirft
- swagger
- grpc


## spec 

this is the spec for the version V1Alpha, it will change almost every time possible nothing that is written here is the last word


### V1Alpha:

tings that should be ready for the first alpha release

 - [V1Alpha.Application](V1Alpha/Application.md)
 - [V1Alpha.Enumeration](V1Alpha/Enumeration.md)
 - [V1Alpha.Enumeration](V1Alpha/EnumerationItem.md)
 - [V1Alpha.Resource](V1Alpha/ResourceStruct.md)
 - [V1Alpha.Resource](V1Alpha/ResourceTable.md)
 - [V1Alpha.Command](V1Alpha/Command.md)
 - [V1Alpha.Query](V1Alpha/Query.md)
 - [V1Alpha.Component](V1Alpha/Component.md)


### V1Next

things that would be  nice to have 

 - [V1Alpha.Branch](V1Next/branch.md)
 - [V1Alpha.Environment](V1Alpha/Environment.md)
 - [V1Alpha.Version](V1Alpha/Version.md)





### Omniql APp
   Omniql application on top of omniql definition that allow to manage resources by federation, orgs , team and users

 - [V1Alpha.Federation](V1Alpha/Federation.md)
 - [V1Alpha.Organization](V1Alpha/Organization.md)
 - [V1Alpha.Team](V1Alpha/Team.md)
 - [V1Alpha.Users](V1Alpha/Users.md)









