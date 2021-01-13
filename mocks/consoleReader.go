/*
Copyright 2019 D2L Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocks

type ConsoleReaderMock struct{}

func (r ConsoleReaderMock) ReadLine(prompt string) (string, error) {
	return prompt, nil
}

func (r ConsoleReaderMock) ReadInt(prompt string) (int, error) {
	return 0, nil
}

func (r ConsoleReaderMock) ReadPassword(prompt string) (string, error) {
	return prompt, nil
}

func (r ConsoleReaderMock) Println(prompt string) error {
	return nil
}

func (r ConsoleReaderMock) Print(prompt string) error {
	return nil
}

func (r ConsoleReaderMock) Option(prompt string, options []string) (int, error) {
	return 0, nil
}
