package tdk

import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
)

func init() {
    fmt.Println("init: tdk");
    android.RegisterModuleType("libteec_go_defaults", libteec_DefaultsFactory)
    android.RegisterModuleType("libckteec_go_defaults", libckteec_DefaultsFactory)
    android.RegisterModuleType("libteec_sys_go_defaults", libteec_sys_DefaultsFactory)
    android.RegisterModuleType("tee_supplicant_go_defaults", tee_supplicant_DefaultsFactory)
    android.RegisterModuleType("tee_stest_go_defaults", tee_stest_DefaultsFactory)
}

func libteec_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        type props struct {
            Enabled *bool
            Proprietary *bool
            Compile_multilib *string
            Multilib struct {
                Lib32 struct {
                    Srcs []string
                    Export_include_dirs []string
                }
                Lib64 struct {
                    Srcs []string
                    Export_include_dirs []string
                }
            }
        }

        p := &props{}
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        tdk_ver := vconfig.String("tdk_version")

        p.Enabled = proptools.BoolPtr(true)
        p.Proprietary = proptools.BoolPtr(true)
        p.Compile_multilib = proptools.StringPtr("both")

        if (tdk_ver == "TDK24") || (tdk_ver == "TDK38") {
            fmt.Println("libteec for", tdk_ver)
            p.Multilib.Lib32.Srcs = []string{"v3.8.0/ca_export_arm/lib_android/libteec.so"}
            p.Multilib.Lib32.Export_include_dirs = []string{"v3.8.0/ca_export_arm/include"}
            p.Multilib.Lib64.Srcs = []string{"v3.8.0/ca_export_arm64/lib_android/libteec.so"}
            p.Multilib.Lib64.Export_include_dirs = []string{"v3.8.0/ca_export_arm64/include"}
        } else if (tdk_ver == "TDK318") {
            fmt.Println("libteec for", tdk_ver)
            p.Multilib.Lib32.Srcs = []string{"v3.18.0/ca_export_arm/lib_android/libteec.so"}
            p.Multilib.Lib32.Export_include_dirs = []string{"v3.18.0/ca_export_arm/include"}
            p.Multilib.Lib64.Srcs = []string{"v3.18.0/ca_export_arm64/lib_android/libteec.so"}
            p.Multilib.Lib64.Export_include_dirs = []string{"v3.18.0/ca_export_arm64/include"}
        } else {
            fmt.Println("No libteec for unknown TDK version")
            p.Enabled = proptools.BoolPtr(false)
        }

        ctx.AppendProperties(p)
    })
    return module
}

func libckteec_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        type props struct {
            Enabled *bool
            Proprietary *bool
            Compile_multilib *string
            Shared_libs []string
            Multilib struct {
                Lib32 struct {
                    Srcs []string
                    Export_include_dirs []string
                }
                Lib64 struct {
                    Srcs []string
                    Export_include_dirs []string
                }
            }
        }

        p := &props{}
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        tdk_ver := vconfig.String("tdk_version")

        p.Enabled = proptools.BoolPtr(true)
        p.Proprietary = proptools.BoolPtr(true)
        p.Compile_multilib = proptools.StringPtr("both")
        p.Shared_libs = []string{"libteec"}

        if (tdk_ver == "TDK24") || (tdk_ver == "TDK38") {
            fmt.Println("No libckteec for", tdk_ver)
            p.Enabled = proptools.BoolPtr(false)
        } else if (tdk_ver == "TDK318") {
            fmt.Println("libckteec for", tdk_ver)
            p.Multilib.Lib32.Srcs = []string{"v3.18.0/ca_export_arm/lib_android/libckteec.so"}
            p.Multilib.Lib32.Export_include_dirs = []string{"v3.18.0/ca_export_arm/include"}
            p.Multilib.Lib64.Srcs = []string{"v3.18.0/ca_export_arm64/lib_android/libckteec.so"}
            p.Multilib.Lib64.Export_include_dirs = []string{"v3.18.0/ca_export_arm64/include"}
        } else {
            fmt.Println("No libckteec for unknown TDK version")
            p.Enabled = proptools.BoolPtr(false)
        }

        ctx.AppendProperties(p)
    })
    return module
}

func libteec_sys_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        type props struct {
            Enabled *bool
            System_ext_specific *bool
            Compile_multilib *string
            Multilib struct {
                Lib32 struct {
                    Srcs []string
                    Export_include_dirs []string
                }
                Lib64 struct {
                    Srcs []string
                    Export_include_dirs []string
                }
            }
        }

        p := &props{}
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        tdk_ver := vconfig.String("tdk_version")

        p.Enabled = proptools.BoolPtr(true)
        p.System_ext_specific = proptools.BoolPtr(true)
        p.Compile_multilib = proptools.StringPtr("both")

        if (tdk_ver == "TDK24") || (tdk_ver == "TDK38") {
            fmt.Println("libteec_sys for", tdk_ver)
            p.Multilib.Lib32.Srcs = []string{"v3.8.0/ca_export_arm/lib_android/libteec_sys.so"}
            p.Multilib.Lib32.Export_include_dirs = []string{"v3.8.0/ca_export_arm/include"}
            p.Multilib.Lib64.Srcs = []string{"v3.8.0/ca_export_arm64/lib_android/libteec_sys.so"}
            p.Multilib.Lib64.Export_include_dirs = []string{"v3.8.0/ca_export_arm64/include"}
        } else if (tdk_ver == "TDK318") {
            fmt.Println("libteec_sys for", tdk_ver)
            p.Multilib.Lib32.Srcs = []string{"v3.18.0/ca_export_arm/lib_android/libteec_sys.so"}
            p.Multilib.Lib32.Export_include_dirs = []string{"v3.18.0/ca_export_arm/include"}
            p.Multilib.Lib64.Srcs = []string{"v3.18.0/ca_export_arm64/lib_android/libteec_sys.so"}
            p.Multilib.Lib64.Export_include_dirs = []string{"v3.18.0/ca_export_arm64/include"}
        } else {
            fmt.Println("No libteec_sys for unknown TDK version")
            p.Enabled = proptools.BoolPtr(false)
        }

        ctx.AppendProperties(p)
    })
    return module
}

func tee_supplicant_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        type props struct {
            Enabled *bool
            Proprietary *bool
            Shared_libs []string
            Init_rc []string
            Arch struct {
                Arm struct {
                    Srcs []string
                }
                Arm64 struct {
                    Srcs []string
                }
            }
        }

        p := &props{}
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        tdk_ver := vconfig.String("tdk_version")

        p.Enabled = proptools.BoolPtr(true)
        p.Proprietary = proptools.BoolPtr(true)
        p.Shared_libs = []string{"libteec"}
        p.Shared_libs = append(p.Shared_libs, "libcutils")

        if (tdk_ver == "TDK24") || (tdk_ver == "TDK38") {
            fmt.Println("tee-supplicant for", tdk_ver)
            p.Arch.Arm.Srcs = []string{"v3.8.0/ca_export_arm/bin_android/tee-supplicant"}
            p.Arch.Arm64.Srcs = []string{"v3.8.0/ca_export_arm64/bin_android/tee-supplicant"}
            p.Init_rc = []string{"v3.8.0/ca_export_arm/bin_android/tee-supplicant.rc"}
        } else if (tdk_ver == "TDK318") {
            fmt.Println("tee-supplicant for", tdk_ver)
            p.Arch.Arm.Srcs = []string{"v3.18.0/ca_export_arm/bin_android/tee-supplicant"}
            p.Arch.Arm64.Srcs = []string{"v3.18.0/ca_export_arm64/bin_android/tee-supplicant"}
            p.Init_rc = []string{"v3.18.0/ca_export_arm/bin_android/tee-supplicant.rc"}
        } else {
            fmt.Println("No tee-supplicant for unknown TDK version")
            p.Enabled = proptools.BoolPtr(false)
        }

        ctx.AppendProperties(p)
    })
    return module
}

func tee_stest_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        type props struct {
            Enabled *bool
            Proprietary *bool
            Shared_libs []string
            Arch struct {
                Arm struct {
                    Srcs []string
                }
                Arm64 struct {
                    Srcs []string
                }
            }
        }

        p := &props{}
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        tdk_ver := vconfig.String("tdk_version")

        if ctx.AConfig().Getenv("TARGET_BUILD_VARIANT") == "userdebug" {
            fmt.Println("Build type is userdebug")
            p.Enabled = proptools.BoolPtr(true)
        } else if ctx.AConfig().Getenv("TARGET_BUILD_VARIANT") == "eng" {
            fmt.Println("Build type is eng")
            p.Enabled = proptools.BoolPtr(true)
        } else if ctx.AConfig().Getenv("TARGET_BUILD_VARIANT") == "user" {
            fmt.Println("Build type is user")
            p.Enabled = proptools.BoolPtr(false)
        } else {
            fmt.Println("Unknown build type")
            p.Enabled = proptools.BoolPtr(false)
        }

        p.Proprietary = proptools.BoolPtr(true)
        p.Shared_libs = []string{"libteec"}

        if (tdk_ver == "TDK24") || (tdk_ver == "TDK38") {
            fmt.Println("tee_stest for", tdk_ver)
            p.Arch.Arm.Srcs = []string{"v3.8.0/ca_export_arm/bin_android/tee_stest"}
            p.Arch.Arm64.Srcs = []string{"v3.8.0/ca_export_arm64/bin_android/tee_stest"}
        } else if (tdk_ver == "TDK318") {
            fmt.Println("tee_stest for", tdk_ver)
            p.Arch.Arm.Srcs = []string{"v3.18.0/ca_export_arm/bin_android/tee_stest"}
            p.Arch.Arm64.Srcs = []string{"v3.18.0/ca_export_arm64/bin_android/tee_stest"}
        } else {
            fmt.Println("No tee_stest for unknown TDK version")
            p.Enabled = proptools.BoolPtr(false)
        }

        ctx.AppendProperties(p)
    })
    return module
}
