# `dllLoader` 包文档

## 概述

`dllLoader` 包为 Go 语言提供了加载和管理本地 DLL 包的功能。通过该包，开发人员可以方便地加载本地 DLL 包，并在应用程序中执行相关操作。

## 结构体

### 1. `DllPackage` 结构体

`DllPackage` 结构体包含了加载的本地 DLL 包的详细信息和操作。

- `name`（字符串）：DLL 包的全局唯一名称。
- `id`（整数）：DLL 包的唯一标识符，由名称和局部唯一的 ID 组成。
- `functions`（字符串切片）：DLL 包支持的函数名列表。
- `info`（映射）：包含其他信息的键值对映射。
- `dll`（`syscall.DLL` 结构体指针）：本地 DLL 包的句柄。

### 2. `DllLoader` 结构体

`DllLoader` 结构体提供了加载和管理本地 DLL 包的功能。

- `Dlls`（嵌套映射）：已加载的 DLL 包的映射，按照 DLL 包名称和 ID 进行组织。
- `dllCounter`（映射）：用于分配 DLL 包 ID 的计数器，按照 DLL 包名称进行组织。

## 功能

### 1. 加载本地 DLL 包

通过 `LoadHexPackage` 方法，可以根据提供的路径加载本地 DLL 包并返回相应的句柄。

### 2. 释放本地 DLL 包

通过 `ReleasePackage` 方法，可以释放加载的本地 DLL 包。

### 3. 获取 DLL 包信息

- `GetName` 方法返回加载的本地 DLL 包的名称。
- `GetID` 方法返回加载的本地 DLL 包的唯一标识符。
- `GetFunctions` 方法返回加载的本地 DLL 包支持的函数名列表。
- `GetInfo` 方法根据提供的键返回加载的本地 DLL 包的附加信息。

## 适用场景

`dllLoader` 包适用于需要在 Go 语言应用程序中加载和管理本地 DLL 包的情景。开发人员可以通过该包轻松实现对 DLL 包的加载、执行和释放等操作，从而扩展应用程序的功能。
