/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package environment

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/mwaa"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/mwaa/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetEnvironmentInput returns input for read
// operation.
func GenerateGetEnvironmentInput(cr *svcapitypes.Environment) *svcsdk.GetEnvironmentInput {
	res := &svcsdk.GetEnvironmentInput{}

	return res
}

// GenerateEnvironment returns the current state in the form of *svcapitypes.Environment.
func GenerateEnvironment(resp *svcsdk.GetEnvironmentOutput) *svcapitypes.Environment {
	cr := &svcapitypes.Environment{}

	if resp.Environment.AirflowConfigurationOptions != nil {
		f0 := map[string]*string{}
		for f0key, f0valiter := range resp.Environment.AirflowConfigurationOptions {
			var f0val string
			f0val = *f0valiter
			f0[f0key] = &f0val
		}
		cr.Spec.ForProvider.AirflowConfigurationOptions = f0
	} else {
		cr.Spec.ForProvider.AirflowConfigurationOptions = nil
	}
	if resp.Environment.AirflowVersion != nil {
		cr.Spec.ForProvider.AirflowVersion = resp.Environment.AirflowVersion
	} else {
		cr.Spec.ForProvider.AirflowVersion = nil
	}
	if resp.Environment.Arn != nil {
		cr.Status.AtProvider.ARN = resp.Environment.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.Environment.DagS3Path != nil {
		cr.Spec.ForProvider.DagS3Path = resp.Environment.DagS3Path
	} else {
		cr.Spec.ForProvider.DagS3Path = nil
	}
	if resp.Environment.EnvironmentClass != nil {
		cr.Spec.ForProvider.EnvironmentClass = resp.Environment.EnvironmentClass
	} else {
		cr.Spec.ForProvider.EnvironmentClass = nil
	}
	if resp.Environment.LastUpdate != nil {
		f8 := &svcapitypes.LastUpdate{}
		if resp.Environment.LastUpdate.Error != nil {
			f8f0 := &svcapitypes.UpdateError{}
			if resp.Environment.LastUpdate.Error.ErrorCode != nil {
				f8f0.ErrorCode = resp.Environment.LastUpdate.Error.ErrorCode
			}
			if resp.Environment.LastUpdate.Error.ErrorMessage != nil {
				f8f0.ErrorMessage = resp.Environment.LastUpdate.Error.ErrorMessage
			}
			f8.Error = f8f0
		}
		if resp.Environment.LastUpdate.Source != nil {
			f8.Source = resp.Environment.LastUpdate.Source
		}
		if resp.Environment.LastUpdate.Status != nil {
			f8.Status = resp.Environment.LastUpdate.Status
		}
		cr.Status.AtProvider.LastUpdate = f8
	} else {
		cr.Status.AtProvider.LastUpdate = nil
	}
	if resp.Environment.LoggingConfiguration != nil {
		f9 := &svcapitypes.LoggingConfigurationInput{}
		if resp.Environment.LoggingConfiguration.DagProcessingLogs != nil {
			f9f0 := &svcapitypes.ModuleLoggingConfigurationInput{}
			if resp.Environment.LoggingConfiguration.DagProcessingLogs.Enabled != nil {
				f9f0.Enabled = resp.Environment.LoggingConfiguration.DagProcessingLogs.Enabled
			}
			if resp.Environment.LoggingConfiguration.DagProcessingLogs.LogLevel != nil {
				f9f0.LogLevel = resp.Environment.LoggingConfiguration.DagProcessingLogs.LogLevel
			}
			f9.DagProcessingLogs = f9f0
		}
		if resp.Environment.LoggingConfiguration.SchedulerLogs != nil {
			f9f1 := &svcapitypes.ModuleLoggingConfigurationInput{}
			if resp.Environment.LoggingConfiguration.SchedulerLogs.Enabled != nil {
				f9f1.Enabled = resp.Environment.LoggingConfiguration.SchedulerLogs.Enabled
			}
			if resp.Environment.LoggingConfiguration.SchedulerLogs.LogLevel != nil {
				f9f1.LogLevel = resp.Environment.LoggingConfiguration.SchedulerLogs.LogLevel
			}
			f9.SchedulerLogs = f9f1
		}
		if resp.Environment.LoggingConfiguration.TaskLogs != nil {
			f9f2 := &svcapitypes.ModuleLoggingConfigurationInput{}
			if resp.Environment.LoggingConfiguration.TaskLogs.Enabled != nil {
				f9f2.Enabled = resp.Environment.LoggingConfiguration.TaskLogs.Enabled
			}
			if resp.Environment.LoggingConfiguration.TaskLogs.LogLevel != nil {
				f9f2.LogLevel = resp.Environment.LoggingConfiguration.TaskLogs.LogLevel
			}
			f9.TaskLogs = f9f2
		}
		if resp.Environment.LoggingConfiguration.WebserverLogs != nil {
			f9f3 := &svcapitypes.ModuleLoggingConfigurationInput{}
			if resp.Environment.LoggingConfiguration.WebserverLogs.Enabled != nil {
				f9f3.Enabled = resp.Environment.LoggingConfiguration.WebserverLogs.Enabled
			}
			if resp.Environment.LoggingConfiguration.WebserverLogs.LogLevel != nil {
				f9f3.LogLevel = resp.Environment.LoggingConfiguration.WebserverLogs.LogLevel
			}
			f9.WebserverLogs = f9f3
		}
		if resp.Environment.LoggingConfiguration.WorkerLogs != nil {
			f9f4 := &svcapitypes.ModuleLoggingConfigurationInput{}
			if resp.Environment.LoggingConfiguration.WorkerLogs.Enabled != nil {
				f9f4.Enabled = resp.Environment.LoggingConfiguration.WorkerLogs.Enabled
			}
			if resp.Environment.LoggingConfiguration.WorkerLogs.LogLevel != nil {
				f9f4.LogLevel = resp.Environment.LoggingConfiguration.WorkerLogs.LogLevel
			}
			f9.WorkerLogs = f9f4
		}
		cr.Spec.ForProvider.LoggingConfiguration = f9
	} else {
		cr.Spec.ForProvider.LoggingConfiguration = nil
	}
	if resp.Environment.MaxWorkers != nil {
		cr.Spec.ForProvider.MaxWorkers = resp.Environment.MaxWorkers
	} else {
		cr.Spec.ForProvider.MaxWorkers = nil
	}
	if resp.Environment.MinWorkers != nil {
		cr.Spec.ForProvider.MinWorkers = resp.Environment.MinWorkers
	} else {
		cr.Spec.ForProvider.MinWorkers = nil
	}
	if resp.Environment.PluginsS3ObjectVersion != nil {
		cr.Spec.ForProvider.PluginsS3ObjectVersion = resp.Environment.PluginsS3ObjectVersion
	} else {
		cr.Spec.ForProvider.PluginsS3ObjectVersion = nil
	}
	if resp.Environment.PluginsS3Path != nil {
		cr.Spec.ForProvider.PluginsS3Path = resp.Environment.PluginsS3Path
	} else {
		cr.Spec.ForProvider.PluginsS3Path = nil
	}
	if resp.Environment.RequirementsS3ObjectVersion != nil {
		cr.Spec.ForProvider.RequirementsS3ObjectVersion = resp.Environment.RequirementsS3ObjectVersion
	} else {
		cr.Spec.ForProvider.RequirementsS3ObjectVersion = nil
	}
	if resp.Environment.RequirementsS3Path != nil {
		cr.Spec.ForProvider.RequirementsS3Path = resp.Environment.RequirementsS3Path
	} else {
		cr.Spec.ForProvider.RequirementsS3Path = nil
	}
	if resp.Environment.Schedulers != nil {
		cr.Spec.ForProvider.Schedulers = resp.Environment.Schedulers
	} else {
		cr.Spec.ForProvider.Schedulers = nil
	}
	if resp.Environment.StartupScriptS3ObjectVersion != nil {
		cr.Spec.ForProvider.StartupScriptS3ObjectVersion = resp.Environment.StartupScriptS3ObjectVersion
	} else {
		cr.Spec.ForProvider.StartupScriptS3ObjectVersion = nil
	}
	if resp.Environment.StartupScriptS3Path != nil {
		cr.Spec.ForProvider.StartupScriptS3Path = resp.Environment.StartupScriptS3Path
	} else {
		cr.Spec.ForProvider.StartupScriptS3Path = nil
	}
	if resp.Environment.Status != nil {
		cr.Status.AtProvider.Status = resp.Environment.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.Environment.Tags != nil {
		f23 := map[string]*string{}
		for f23key, f23valiter := range resp.Environment.Tags {
			var f23val string
			f23val = *f23valiter
			f23[f23key] = &f23val
		}
		cr.Spec.ForProvider.Tags = f23
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.Environment.WebserverAccessMode != nil {
		cr.Spec.ForProvider.WebserverAccessMode = resp.Environment.WebserverAccessMode
	} else {
		cr.Spec.ForProvider.WebserverAccessMode = nil
	}
	if resp.Environment.WeeklyMaintenanceWindowStart != nil {
		cr.Spec.ForProvider.WeeklyMaintenanceWindowStart = resp.Environment.WeeklyMaintenanceWindowStart
	} else {
		cr.Spec.ForProvider.WeeklyMaintenanceWindowStart = nil
	}

	return cr
}

// GenerateCreateEnvironmentInput returns a create input.
func GenerateCreateEnvironmentInput(cr *svcapitypes.Environment) *svcsdk.CreateEnvironmentInput {
	res := &svcsdk.CreateEnvironmentInput{}

	if cr.Spec.ForProvider.AirflowConfigurationOptions != nil {
		f0 := map[string]*string{}
		for f0key, f0valiter := range cr.Spec.ForProvider.AirflowConfigurationOptions {
			var f0val string
			f0val = *f0valiter
			f0[f0key] = &f0val
		}
		res.SetAirflowConfigurationOptions(f0)
	}
	if cr.Spec.ForProvider.AirflowVersion != nil {
		res.SetAirflowVersion(*cr.Spec.ForProvider.AirflowVersion)
	}
	if cr.Spec.ForProvider.DagS3Path != nil {
		res.SetDagS3Path(*cr.Spec.ForProvider.DagS3Path)
	}
	if cr.Spec.ForProvider.EnvironmentClass != nil {
		res.SetEnvironmentClass(*cr.Spec.ForProvider.EnvironmentClass)
	}
	if cr.Spec.ForProvider.LoggingConfiguration != nil {
		f4 := &svcsdk.LoggingConfigurationInput{}
		if cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs != nil {
			f4f0 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.Enabled != nil {
				f4f0.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.LogLevel != nil {
				f4f0.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.LogLevel)
			}
			f4.SetDagProcessingLogs(f4f0)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs != nil {
			f4f1 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.Enabled != nil {
				f4f1.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.LogLevel != nil {
				f4f1.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.LogLevel)
			}
			f4.SetSchedulerLogs(f4f1)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.TaskLogs != nil {
			f4f2 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.Enabled != nil {
				f4f2.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.LogLevel != nil {
				f4f2.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.LogLevel)
			}
			f4.SetTaskLogs(f4f2)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs != nil {
			f4f3 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.Enabled != nil {
				f4f3.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.LogLevel != nil {
				f4f3.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.LogLevel)
			}
			f4.SetWebserverLogs(f4f3)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs != nil {
			f4f4 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.Enabled != nil {
				f4f4.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.LogLevel != nil {
				f4f4.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.LogLevel)
			}
			f4.SetWorkerLogs(f4f4)
		}
		res.SetLoggingConfiguration(f4)
	}
	if cr.Spec.ForProvider.MaxWorkers != nil {
		res.SetMaxWorkers(*cr.Spec.ForProvider.MaxWorkers)
	}
	if cr.Spec.ForProvider.MinWorkers != nil {
		res.SetMinWorkers(*cr.Spec.ForProvider.MinWorkers)
	}
	if cr.Spec.ForProvider.PluginsS3ObjectVersion != nil {
		res.SetPluginsS3ObjectVersion(*cr.Spec.ForProvider.PluginsS3ObjectVersion)
	}
	if cr.Spec.ForProvider.PluginsS3Path != nil {
		res.SetPluginsS3Path(*cr.Spec.ForProvider.PluginsS3Path)
	}
	if cr.Spec.ForProvider.RequirementsS3ObjectVersion != nil {
		res.SetRequirementsS3ObjectVersion(*cr.Spec.ForProvider.RequirementsS3ObjectVersion)
	}
	if cr.Spec.ForProvider.RequirementsS3Path != nil {
		res.SetRequirementsS3Path(*cr.Spec.ForProvider.RequirementsS3Path)
	}
	if cr.Spec.ForProvider.Schedulers != nil {
		res.SetSchedulers(*cr.Spec.ForProvider.Schedulers)
	}
	if cr.Spec.ForProvider.StartupScriptS3ObjectVersion != nil {
		res.SetStartupScriptS3ObjectVersion(*cr.Spec.ForProvider.StartupScriptS3ObjectVersion)
	}
	if cr.Spec.ForProvider.StartupScriptS3Path != nil {
		res.SetStartupScriptS3Path(*cr.Spec.ForProvider.StartupScriptS3Path)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range cr.Spec.ForProvider.Tags {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		res.SetTags(f14)
	}
	if cr.Spec.ForProvider.WebserverAccessMode != nil {
		res.SetWebserverAccessMode(*cr.Spec.ForProvider.WebserverAccessMode)
	}
	if cr.Spec.ForProvider.WeeklyMaintenanceWindowStart != nil {
		res.SetWeeklyMaintenanceWindowStart(*cr.Spec.ForProvider.WeeklyMaintenanceWindowStart)
	}

	return res
}

// GenerateUpdateEnvironmentInput returns an update input.
func GenerateUpdateEnvironmentInput(cr *svcapitypes.Environment) *svcsdk.UpdateEnvironmentInput {
	res := &svcsdk.UpdateEnvironmentInput{}

	if cr.Spec.ForProvider.AirflowConfigurationOptions != nil {
		f0 := map[string]*string{}
		for f0key, f0valiter := range cr.Spec.ForProvider.AirflowConfigurationOptions {
			var f0val string
			f0val = *f0valiter
			f0[f0key] = &f0val
		}
		res.SetAirflowConfigurationOptions(f0)
	}
	if cr.Spec.ForProvider.AirflowVersion != nil {
		res.SetAirflowVersion(*cr.Spec.ForProvider.AirflowVersion)
	}
	if cr.Spec.ForProvider.DagS3Path != nil {
		res.SetDagS3Path(*cr.Spec.ForProvider.DagS3Path)
	}
	if cr.Spec.ForProvider.EnvironmentClass != nil {
		res.SetEnvironmentClass(*cr.Spec.ForProvider.EnvironmentClass)
	}
	if cr.Spec.ForProvider.LoggingConfiguration != nil {
		f4 := &svcsdk.LoggingConfigurationInput{}
		if cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs != nil {
			f4f0 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.Enabled != nil {
				f4f0.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.LogLevel != nil {
				f4f0.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.DagProcessingLogs.LogLevel)
			}
			f4.SetDagProcessingLogs(f4f0)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs != nil {
			f4f1 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.Enabled != nil {
				f4f1.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.LogLevel != nil {
				f4f1.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.SchedulerLogs.LogLevel)
			}
			f4.SetSchedulerLogs(f4f1)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.TaskLogs != nil {
			f4f2 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.Enabled != nil {
				f4f2.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.LogLevel != nil {
				f4f2.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.TaskLogs.LogLevel)
			}
			f4.SetTaskLogs(f4f2)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs != nil {
			f4f3 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.Enabled != nil {
				f4f3.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.LogLevel != nil {
				f4f3.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.WebserverLogs.LogLevel)
			}
			f4.SetWebserverLogs(f4f3)
		}
		if cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs != nil {
			f4f4 := &svcsdk.ModuleLoggingConfigurationInput{}
			if cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.Enabled != nil {
				f4f4.SetEnabled(*cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.Enabled)
			}
			if cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.LogLevel != nil {
				f4f4.SetLogLevel(*cr.Spec.ForProvider.LoggingConfiguration.WorkerLogs.LogLevel)
			}
			f4.SetWorkerLogs(f4f4)
		}
		res.SetLoggingConfiguration(f4)
	}
	if cr.Spec.ForProvider.MaxWorkers != nil {
		res.SetMaxWorkers(*cr.Spec.ForProvider.MaxWorkers)
	}
	if cr.Spec.ForProvider.MinWorkers != nil {
		res.SetMinWorkers(*cr.Spec.ForProvider.MinWorkers)
	}
	if cr.Spec.ForProvider.PluginsS3ObjectVersion != nil {
		res.SetPluginsS3ObjectVersion(*cr.Spec.ForProvider.PluginsS3ObjectVersion)
	}
	if cr.Spec.ForProvider.PluginsS3Path != nil {
		res.SetPluginsS3Path(*cr.Spec.ForProvider.PluginsS3Path)
	}
	if cr.Spec.ForProvider.RequirementsS3ObjectVersion != nil {
		res.SetRequirementsS3ObjectVersion(*cr.Spec.ForProvider.RequirementsS3ObjectVersion)
	}
	if cr.Spec.ForProvider.RequirementsS3Path != nil {
		res.SetRequirementsS3Path(*cr.Spec.ForProvider.RequirementsS3Path)
	}
	if cr.Spec.ForProvider.Schedulers != nil {
		res.SetSchedulers(*cr.Spec.ForProvider.Schedulers)
	}
	if cr.Spec.ForProvider.StartupScriptS3ObjectVersion != nil {
		res.SetStartupScriptS3ObjectVersion(*cr.Spec.ForProvider.StartupScriptS3ObjectVersion)
	}
	if cr.Spec.ForProvider.StartupScriptS3Path != nil {
		res.SetStartupScriptS3Path(*cr.Spec.ForProvider.StartupScriptS3Path)
	}
	if cr.Spec.ForProvider.WebserverAccessMode != nil {
		res.SetWebserverAccessMode(*cr.Spec.ForProvider.WebserverAccessMode)
	}
	if cr.Spec.ForProvider.WeeklyMaintenanceWindowStart != nil {
		res.SetWeeklyMaintenanceWindowStart(*cr.Spec.ForProvider.WeeklyMaintenanceWindowStart)
	}

	return res
}

// GenerateDeleteEnvironmentInput returns a deletion input.
func GenerateDeleteEnvironmentInput(cr *svcapitypes.Environment) *svcsdk.DeleteEnvironmentInput {
	res := &svcsdk.DeleteEnvironmentInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}
