package influxPlay

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go/domain"
	"github.com/joho/godotenv"
)

func Test_connectToInfluxDB(t *testing.T) {

	//load environment variable from a file for test purposes
	godotenv.Load("./env.env")

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful connection to InfluxDB",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectToInfluxDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToInfluxDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			health, err := got.Health(context.Background())
			if (err != nil) && health.Status == domain.HealthCheckStatusPass {
				t.Errorf("connectToInfluxDB() error. database not healthy")
				return
			}
			got.Close()
		})
	}
}

func Test_WriteLineToInfluxDB(t *testing.T) {

	//load environment variable from a file for test purposes
	godotenv.Load("./env.env")

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Written data",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ConnectToInfluxDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToInfluxDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := 0; i < 100; i++ {
				line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", (rand.Float64()*5)+10, (rand.Float64()*5)+15)
				WriteLineToInfluxDB(line, client, "playground", "iot")
				fmt.Println(line)
				time.Sleep(100 * time.Millisecond)
			}

		})
	}
}
