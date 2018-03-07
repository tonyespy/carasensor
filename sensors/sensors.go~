// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2018 Canonical Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package sensors

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"	
	"path/filepath"
	"strconv"
	"strings"
)

const debug = true
const iioDevices = "/sys/bus/iio/devices/"

type Sensor interface {
	init()
	readSensorAttr(attr string) (float64, error)
}

type Lng2dm interface {
	AccelX() (float64, error)
	AccelY() (float64, error)
	AccelZ() (float64, error)
}

type Hts221 interface {
	Humidity() (float64, error)
	Temperature() (float64, error)
}

type device struct {
	name string
	path string
}

func (d *device) init() {
	files, err := ioutil.ReadDir(iioDevices)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		path := filepath.Join(iioDevices, f.Name())
		nameFile := filepath.Join(path, "name")

		dat, err := ioutil.ReadFile(nameFile)
		if err == nil {
			if strings.TrimSpace(string(dat)) == d.name {
				fmt.Println(string(dat))				
				d.path = path
			}
		}
	}
}

func (d *device) readSensorAttr(attr string) (float64, error) {
	path := filepath.Join(d.path, attr)

	f, err := os.Open(path)
	if err != nil {
		log.Println("Can't open sensor attribute: %v", err);
	}

	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Println("Can't read sensor attribute: %v", err);
	}

	val, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		log.Println("Can't convert sensor attribute: %v", err);
	}
	
	return val, nil
}

// New returns a Systemd that uses the given rootDir
func NewLng2dm() Lng2dm {
	d := &device{name: "lng2dm"}
	d.init()
	return d
}

// New returns a Systemd that uses the given rootDir
func NewHts221() Hts221 {
	d := &device{name: "hts221"}
	d.init()
	return d
}

func (d *device) AccelX() (float64, error) {
	raw, err := d.readSensorAttr("in_accel_x_raw")
	if err != nil {
		return 0.0, err
	}
	
	scale, err := d.readSensorAttr("in_accel_x_scale")
	if err != nil {
		return 0.0, err
	}
	
	if debug == true {
		log.Printf("x_raw: %f x_scale: %f\n", raw, scale)
	}
	
	return raw * scale, nil
}

func (d *device) AccelY() (float64, error) {
	raw, err := d.readSensorAttr("in_accel_y_raw")
	if err != nil {
		return 0.0, err
	}
	
	scale, err := d.readSensorAttr("in_accel_y_scale")
	if err != nil {
		return 0.0, err
	}
	
	if debug == true {
		log.Printf("y_raw: %f y_scale: %f\n", raw, scale)
	}
	
	return raw * scale, nil
}

func (d *device) AccelZ() (float64, error) {
	raw, err := d.readSensorAttr("in_accel_z_raw")
	if err != nil {
		return 0.0, err
	}
	
	scale, err := d.readSensorAttr("in_accel_z_scale")
	if err != nil {
		return 0.0, err
	}
	
	if debug == true {
		log.Printf("z_raw: %f z_scale: %f\n", raw, scale)
	}
	
	return raw * scale, nil
}

func (d *device) Humidity() (float64, error) {
	raw, err := d.readSensorAttr("in_humidityrelative_raw")
	if err != nil {
		return 0.0, err
	}

	offset, err := d.readSensorAttr("in_humidityrelative_offset")
	if err != nil {
		return 0.0, err
	}

	scale, err := d.readSensorAttr("in_humidityrelative_scale")
	if err != nil {
		return 0.0, err
	}
	
	if debug == true {
		log.Printf("raw: %f offset: %f scale: %f\n", raw, offset, scale)
	}
	
	return (raw + offset) * scale, nil
}

func (d *device) Temperature() (float64, error) {
	raw, err := d.readSensorAttr("in_temp_raw")
	if err != nil {
		return 0.0, err
	}

	offset, err := d.readSensorAttr("in_temp_offset")
	if err != nil {
		return 0.0, err
	}

	scale, err := d.readSensorAttr("in_temp_scale")
	if err != nil {
		return 0.0, err
	}
	
	if debug == true {
		log.Printf("raw: %f offset: %f scale: %f\n", raw, offset, scale)
	}
	
	return (raw + offset) * scale, nil
}
