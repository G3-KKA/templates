__Usage guide__

# To modify config
 Use config.go as API 
 Just throw everyting you need into specified places:
 - ENV
 - flag binds
 - config file structure
 - other viper options
 - force overrides

Everything will work automaticaly
# In-Program Usage
 config.InitConfig should be called at the start
 config.Get will return last version of config

Thats it!  

See ../example/example.go for hints 
And repeat 

You may also want to visit https://zhwt.github.io/yaml-to-go/
it generates go struct from your .yaml

Then replace all `yaml:"..."` to   `mapstructure:"..."`
This is essential to work with viper library
That template based on it 