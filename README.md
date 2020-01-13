# drone-sonar-scan-plugin

Drone plugin to run sonarqube scanner

Alternative to [aosapps/drone-sonar-plugin](https://github.com/aosapps/drone-sonar-plugin)

# usage

Drone step:
```yml
steps
- name: sonar scan
  image: fmetais/drone-sonar-scan-plugin:latest  
  settings:
    sonar.host.url: 
      from_secret: sonar_host
    sonar.login: 
      from_secret: sonar_token
```

Sonar properties are taken from sonar-project.properties.
Your can overide this param with drone settings:

Drone step:
```yml
steps
- name: sonar scan
  image: fmetais/drone-sonar-scan-plugin:latest
  settings:
    sonar.host.url: 
      from_secret: sonar_host
    sonar.login: 
      from_secret: sonar_token
    sonar.project-key: my-key
    sonar.project-name: My project
    sonar.log.level: debug
    sonar.show-profiling: true
    ....

```
Drone settings are passed to sonar-scanner command.
([)Sonar scanner documentation) [https://docs.sonarqube.org/latest/analysis/analysis-parameters/].