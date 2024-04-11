# MNC CLI

> Tool to manage your Minecraft servers/clients easily

## Manifest

1. Манифест представляет собой один или несколько YAML файлов с документами.
2. Каждый документ в K8s подобной манере описывает ресурс.
   1. Ресурс имеет свой тип и версию.
3. Ресурсы могут быть зависимыми от других ресурсов по типу.
4. Ресурсы могут создавать новые виртуальные ресурсы.
   1. Ресурс периодически проверяет условия своей готовности (проверяет зависимости).
   2. Ресурс периодически проверяет валидность зависимых ресурсов (пересоздаёт их по необходимости).

Для начала нам нужно только 2 вида манифестов: asset/remote и minecraft/server/mohist чтобы описать
существующую инфраструктуру.

```yaml
# asset/remote v1 - позволяет описать один удалённый файл для последующей работы.
apiVersion: v1
kind: asset/remote
metadata:
    name: AppleSkin_Forge_1.20.1 # inner MNC identity name
    labels: # selectable fields
        minecraftLoader: Forge # Forge, NeoForge, Fabric, Paper, Spigot, etc.
        minecraftVersion: 1.20.1
        minecraftKind: Mod # Mod, Plugin, Shader, DataPack, ResourcePack, ModPack, etc.
        minecraftSide: Client # Server, Client, Both
        serverStone: true # user defined label to select this resource by.
    annotations: # optional/informative fields
        assetTitle: AppleSkin
        assetDescription: >
            Adds various food-related HUD improvements. This is a mostly client-side mod, but it needs to be on the
            server in order to display accurate saturation/exhaustion values on the client.
        assetLink: https://modrinth.com/mod/appleskin
        assetIcon: https://cdn.modrinth.com/data/EsAfCjCV/icon.png
        assetAuthorName: squeek502
        assetAuthorLink: https://modrinth.com/user/squeek502
spec:
    pullPolicy: IfNotPresent
    url: https://cdn.modrinth.com/data/EsAfCjCV/versions/XdXDExVF/appleskin-forge-mc1.20.1-2.5.1.jar

---

# Позволяет генерировать набор asset/remote/v1 из списка url'ов.
# ! Временно не позволяем подобный набор ресурсов, который может порождать другие ресурсы.
# ! Необходимо детально продумать механизм порождения ресурсов и статуса их готовности.
apiVersion: v1
kind: asset/remote/list
metadata:
    name: Mods_Server_Stone_Forge_1.20.1
    labels:
        minecraftLoader: forge
        minecraftVersion: 1.20.1
spec:
    template:
        metadata:
            name: Mod_Server_Stone_Forge_1.20.1_
            labels:
                minecraftLoader: forge
                minecraftVersion: 1.20.1
                serverStone: true
        spec:
            pullPolicy: IfNotPresent
    urls:
        - https://cdn.modrinth.com/data/EsAfCjCV/versions/XdXDExVF/appleskin-forge-mc1.20.1-2.5.1.jar

---

# Позволяет описать набор ассетов используя лейблы как селекторы.
apiVersion: v1
kind: asset/pack
metadata:
    name: MNCRED_Server_Stone_Server_Mods
    labels:
        serverStone: true
spec:
    selector:
        matchLabels:
            serverStone: true

---

apiVersion: v1
kind: minecraft/server/mohist
metadata:
    name: MNCRED_Server_Stone
spec:
    #TODO: Добавить информацию по которой возможно скачать и установить mohist сервер.
    assets:
        # accepts multiple assets and put it all inside dir.
        - target: mods/ # the suffix slash is required to make it a directory. File allowed for only single selected resource or panic.
          selector:
            matchLabels:
                serverStone: true
                serverMod: true
        # accepts only one asset with server.properties content and put it as a file. Panics if many resources found.
        - target: server.properties
          selector:
            matchLabels:
                serverStone: true
                serverProperties: true

```

Может помочь подобный проект на Go: [https://github.com/nothub/mrpack-install]
