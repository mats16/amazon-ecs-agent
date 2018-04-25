// +build !integration
// Copyright 2014-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package dockerstate

import (
	"testing"

	"github.com/aws/amazon-ecs-agent/agent/api"
	apicontainer "github.com/aws/amazon-ecs-agent/agent/api/container"
	"github.com/aws/amazon-ecs-agent/agent/engine/image"

	"github.com/stretchr/testify/assert"
)

func TestCreateDockerTaskEngineState(t *testing.T) {
	state := NewTaskEngineState()

	if _, ok := state.ContainerByID("test"); ok {
		t.Error("Empty state should not have a test container")
	}

	if _, ok := state.ContainerMapByArn("test"); ok {
		t.Error("Empty state should not have a test task")
	}

	if _, ok := state.TaskByShortID("test"); ok {
		t.Error("Empty state should not have a test taskid")
	}

	if _, ok := state.TaskByID("test"); ok {
		t.Error("Empty state should not have a test taskid")
	}

	if len(state.AllTasks()) != 0 {
		t.Error("Empty state should have no tasks")
	}

	if len(state.AllImageStates()) != 0 {
		t.Error("Empty state should have no image states")
	}

	assert.Len(t, state.(*DockerTaskEngineState).AllENIAttachments(), 0)
	task, ok := state.TaskByShortID("test")
	if assert.Empty(t, ok, "Empty state should have no tasks") {
		assert.Empty(t, task, "Empty state should have no tasks")
	}

	assert.Empty(t, state.GetAllContainerIDs(), "Empty state should have no containers")
}

func TestAddTask(t *testing.T) {
	state := NewTaskEngineState()

	testTask := &api.Task{Arn: "test"}
	state.AddTask(testTask)

	if len(state.AllTasks()) != 1 {
		t.Error("Should have 1 task")
	}

	task, ok := state.TaskByArn("test")
	if !ok {
		t.Error("Couldn't find the test task")
	}
	if task.Arn != "test" {
		t.Error("Wrong task retrieved")
	}
}

func TestAddRemoveENIAttachment(t *testing.T) {
	state := NewTaskEngineState()

	attachment := &api.ENIAttachment{
		TaskARN:       "taskarn",
		AttachmentARN: "eni1",
		MACAddress:    "mac1",
	}

	state.AddENIAttachment(attachment)
	assert.Len(t, state.(*DockerTaskEngineState).AllENIAttachments(), 1)
	eni, ok := state.ENIByMac("mac1")
	assert.True(t, ok)
	assert.Equal(t, eni.TaskARN, attachment.TaskARN)

	eni, ok = state.ENIByMac("non-mac")
	assert.False(t, ok)
	assert.Nil(t, eni)

	// Remove the attachment from state
	state.RemoveENIAttachment(attachment.MACAddress)
	assert.Len(t, state.AllImageStates(), 0)
	eni, ok = state.ENIByMac("mac1")
	assert.False(t, ok)
	assert.Nil(t, eni)
}

func TestTwophaseAddContainer(t *testing.T) {
	state := NewTaskEngineState()
	testTask := &api.Task{Arn: "test", Containers: []*apicontainer.Container{{
		Name: "testContainer",
	}}}
	state.AddTask(testTask)

	state.AddContainer(&apicontainer.DockerContainer{DockerName: "dockerName", Container: testTask.Containers[0]}, testTask)

	if len(state.AllTasks()) != 1 {
		t.Fatal("Should have 1 task")
	}

	task, ok := state.TaskByArn("test")
	if !ok {
		t.Error("Couldn't find the test task")
	}
	if task.Arn != "test" {
		t.Error("Wrong task retrieved")
	}

	containerMap, ok := state.ContainerMapByArn("test")
	if !ok {
		t.Fatal("Could not get container map")
	}

	container, ok := containerMap["testContainer"]
	if !ok {
		t.Fatal("Could not get container")
	}
	if container.DockerName != "dockerName" {
		t.Fatal("Incorrect docker name")
	}
	if container.DockerID != "" {
		t.Fatal("DockerID Should be blank")
	}

	state.AddContainer(&apicontainer.DockerContainer{DockerName: "dockerName", Container: testTask.Containers[0], DockerID: "did"}, testTask)

	containerMap, ok = state.ContainerMapByArn("test")
	if !ok {
		t.Fatal("Could not get container map")
	}

	container, ok = containerMap["testContainer"]
	if !ok {
		t.Fatal("Could not get container")
	}
	if container.DockerName != "dockerName" {
		t.Fatal("Incorrect docker name")
	}
	if container.DockerID != "did" {
		t.Fatal("DockerID should have been updated")
	}

	container, ok = state.ContainerByID("did")
	if !ok {
		t.Fatal("Could not get container by id")
	}
	if container.DockerName != "dockerName" || container.DockerID != "did" {
		t.Fatal("Incorrect container fetched")
	}
}

func TestRemoveTask(t *testing.T) {
	state := NewTaskEngineState()
	testContainer1 := &apicontainer.Container{
		Name: "c1",
	}

	containerID := "did"
	testDockerContainer1 := &apicontainer.DockerContainer{
		DockerID:  containerID,
		Container: testContainer1,
	}
	testContainer2 := &apicontainer.Container{
		Name: "c2",
	}
	testDockerContainer2 := &apicontainer.DockerContainer{
		// DockerName is used before the DockerID is assigned
		DockerName: "docker-name-2",
		Container:  testContainer2,
	}
	testTask := &api.Task{
		Arn:        "t1",
		Containers: []*apicontainer.Container{testContainer1, testContainer2},
	}

	state.AddTask(testTask)
	state.AddContainer(testDockerContainer1, testTask)
	state.AddContainer(testDockerContainer2, testTask)
	addr := "169.254.170.3"
	state.AddTaskIPAddress(addr, testTask.Arn)
	engineState := state.(*DockerTaskEngineState)

	assert.Len(t, state.AllTasks(), 1, "Expected one task")
	assert.Len(t, engineState.idToTask, 2, "idToTask map should have two entries")
	assert.Len(t, engineState.idToContainer, 2, "idToContainer map should have two entries")
	taskARNFromIP, ok := state.GetTaskByIPAddress(addr)
	assert.True(t, ok)
	assert.Equal(t, testTask.Arn, taskARNFromIP)

	state.RemoveTask(testTask)

	assert.Len(t, state.AllTasks(), 0, "Expected task to be removed")
	assert.Len(t, engineState.idToTask, 0, "idToTask map should be empty")
	assert.Len(t, engineState.idToContainer, 0, "idToContainer map should be empty")
	_, ok = state.GetTaskByIPAddress(addr)
	assert.False(t, ok)
}

func TestAddImageState(t *testing.T) {
	state := NewTaskEngineState()

	testImage := &image.Image{ImageID: "sha256:imagedigest"}
	testImageState := &image.ImageState{Image: testImage}
	state.AddImageState(testImageState)

	if len(state.AllImageStates()) != 1 {
		t.Error("Error adding image state")
	}

	for _, imageState := range state.AllImageStates() {
		if imageState.Image.ImageID != testImage.ImageID {
			t.Error("Error in retrieving image state added")
		}
	}
}

func TestAddEmptyImageState(t *testing.T) {
	state := NewTaskEngineState()
	state.AddImageState(nil)

	if len(state.AllImageStates()) != 0 {
		t.Error("Error adding empty image state")
	}
}

func TestAddEmptyIdImageState(t *testing.T) {
	state := NewTaskEngineState()

	testImage := &image.Image{ImageID: ""}
	testImageState := &image.ImageState{Image: testImage}
	state.AddImageState(testImageState)

	if len(state.AllImageStates()) != 0 {
		t.Error("Error adding image state with empty Image Id")
	}
}

func TestRemoveImageState(t *testing.T) {
	state := NewTaskEngineState()

	testImage := &image.Image{ImageID: "sha256:imagedigest"}
	testImageState := &image.ImageState{Image: testImage}
	state.AddImageState(testImageState)

	if len(state.AllImageStates()) != 1 {
		t.Error("Error adding image state")
	}
	state.RemoveImageState(testImageState)
	if len(state.AllImageStates()) != 0 {
		t.Error("Error removing image state")
	}
}

func TestRemoveEmptyImageState(t *testing.T) {
	state := NewTaskEngineState()

	testImage := &image.Image{ImageID: "sha256:imagedigest"}
	testImageState := &image.ImageState{Image: testImage}
	state.AddImageState(testImageState)

	if len(state.AllImageStates()) != 1 {
		t.Error("Error adding image state")
	}
	state.RemoveImageState(nil)
	if len(state.AllImageStates()) == 0 {
		t.Error("Error removing empty image state")
	}
}

func TestRemoveNonExistingImageState(t *testing.T) {
	state := NewTaskEngineState()

	testImage := &image.Image{ImageID: "sha256:imagedigest"}
	testImageState := &image.ImageState{Image: testImage}
	state.AddImageState(testImageState)

	if len(state.AllImageStates()) != 1 {
		t.Error("Error adding image state")
	}
	testImage1 := &image.Image{ImageID: "sha256:imagedigest1"}
	testImageState1 := &image.ImageState{Image: testImage1}
	state.RemoveImageState(testImageState1)
	if len(state.AllImageStates()) == 0 {
		t.Error("Error removing incorrect image state")
	}
}

// TestAddContainer tests first add container with docker name and
// then add the container with dockerID
func TestAddContainerNameAndID(t *testing.T) {
	state := NewTaskEngineState()

	task := &api.Task{
		Arn: "taskArn",
	}
	container := &apicontainer.DockerContainer{
		DockerName: "ecs-test-container-1",
		Container: &apicontainer.Container{
			Name: "test",
		},
	}
	state.AddTask(task)
	state.AddContainer(container, task)
	containerMap, ok := state.ContainerMapByArn(task.Arn)
	assert.True(t, ok)
	assert.Len(t, containerMap, 1)

	assert.Len(t, state.GetAllContainerIDs(), 1)

	_, ok = state.ContainerByID(container.DockerName)
	assert.True(t, ok, "container with DockerName should be added to the state")

	container = &apicontainer.DockerContainer{
		DockerName: "ecs-test-container-1",
		DockerID:   "dockerid",
		Container: &apicontainer.Container{
			Name: "test",
		},
	}
	state.AddContainer(container, task)
	assert.Len(t, containerMap, 1)
	assert.Len(t, state.GetAllContainerIDs(), 1)
	_, ok = state.ContainerByID(container.DockerID)
	assert.True(t, ok, "container with DockerName should be added to the state")
	_, ok = state.ContainerByID(container.DockerName)
	assert.False(t, ok, "container with DockerName should be added to the state")
}

func TestTaskIPAddress(t *testing.T) {
	state := newDockerTaskEngineState()
	addr := "169.254.170.3"
	taskARN := "t1"
	state.AddTaskIPAddress(addr, taskARN)
	taskARNFromIP, ok := state.GetTaskByIPAddress(addr)
	assert.True(t, ok)
	assert.Equal(t, taskARN, taskARNFromIP)
	taskIP, ok := state.taskToIPUnsafe(taskARN)
	assert.True(t, ok)
	assert.Equal(t, addr, taskIP)
}
