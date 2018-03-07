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

package main

import (
	"fmt"
	"os"
	"github.com/tonyespy/caracalla/sensors"
)

func main() {	
	if len(os.Args) < 2 {
		fmt.Println("Please specify sensor name (lng2dm, lps22hb, or hts221)")
		return
	}

	switch os.Args[1] {
	case "lng2dm":
		sensor := sensors.NewLng2dm()
		val, err := sensor.AccelX()
		if err != nil {
			fmt.Printf("error reading accel_x: %v\n", err)
			return
		}

		fmt.Printf("accel_x(m/s^): %f\n", val)

		val, err = sensor.AccelY()
		if err != nil {
			fmt.Printf("error reading accel_y: %v\n", err)
			return
		}

		fmt.Printf("accel_y(m/s^): %f\n", val)

		val, err = sensor.AccelZ()
		if err != nil {
			fmt.Println("error reading accel_z: %v", err)
			return
		}
		
		fmt.Printf("accel_z(m/s^): %f\n", val)
		
	case "lps22hb":
		fmt.Println("lps22hb")		
	case "hts221":
		sensor := sensors.NewHts221()
		
		val, err := sensor.Humidity()
		if err != nil {
			fmt.Printf("error reading hts221 humidity: %v\n", err)
			return
		}

		fmt.Printf("hts221: Relative Humidity(%): %f\n", val)

		val, err = sensor.Temperature()
		if err != nil {
			fmt.Printf("error reading hts221 temperature: %v\n", err)
			return
		}

		fmt.Printf("hts221: Temperature(degC): %f\n", val)
	default:
		fmt.Println("Please specify sensor name (lng2dm, lps22hb, or hts221)")
	}


}



