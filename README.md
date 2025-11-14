# 👻 GhostWA — Monitor Local de Mensajes de WhatsApp

GhostWA es una herramienta de monitoreo **local**, diseñada para analizar en tiempo real la base de datos de WhatsApp en dispositivos Android con **acceso root**.  
Permite detectar:

- 📥 Mensajes entrantes  
- ❌ Mensajes eliminados  
- ✏️ Mensajes editados  

**GhostWA NO modifica WhatsApp, NO descifra mensajes y NO se conecta a servidores.**  
Solo lee información local ya almacenada por el propio sistema de WhatsApp.

---

## 🚀 Características principales

- 🔍 **Lectura en tiempo real** de `msgstore.db`
- 👤 Resolución automática de nombres de contactos via `wa.db`
- 🍃 Interfaz CLI con colores y comandos internos
- 🛡️ Modo seguro (protege funciones peligrosas)
- 🛠️ Modo desarrollador (debug avanzado)
- ⏱️ Cambiar intervalo de lectura (`set-poll`)
- 🎛️ Filtrado de tipos de mensajes (`ignore`)
- 🧽 Limpieza, cierre, ayuda y más comandos interactivos

---

## 📌 Requisitos

- Dispositivo Android **rooteado**
- Termux
- Acceso a:
  - `/data/data/com.whatsapp/databases/msgstore.db`
  - `/data/data/com.whatsapp/databases/wa.db`
- SQLite3 instalado

---

## 📥 Instalación

1. Clonar repositorio:

```bash
git clone https://github.com/angelunc/GhostWA.git
cd GhostWA
