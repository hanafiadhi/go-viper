package go_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T)  {
	var config *viper.Viper = viper.New()
	assert.NotNil(t,config)
}

func TestJson(t *testing.T)  {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	assert.Nil(t,err)

	assert.Equal(t, "go-viper", config.GetString("app.name"))
	assert.Equal(t, true, config.GetBool("database.debug"))
	assert.Equal(t, 3000, config.GetInt("database.port"))
}

func TestEnv(t *testing.T)  {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	assert.Nil(t,err)

	assert.Equal(t, "go-viper", config.GetString("APP_NAME"))
	assert.Equal(t, true, config.GetBool("DATABASE_DEBUG"))
	assert.Equal(t, 3000, config.GetInt("DATABSE_PORT"))
}
