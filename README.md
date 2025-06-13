
## Commandos Prisma

Se debe crear el .env por las dudas, aunque en este caso ya esta definida la ruta desde el dockerfile

### Generar Modelo Prisma

```bash
go run github.com/steebchen/prisma-client-go generate
```
### Leer BD con prisma

```bash
go run github.com/steebchen/prisma-client-go db pull
```
### Sincronizar BD con prisma

```bash
go run github.com/steebchen/prisma-client-go db push
```
### Migracion Prisma Client Go 
```bash
Usar con npx riesgo
npx prisma migrate dev --name first_migration
```
### Abrir visor de BD con prisma

```bash
npx prisma studio
```

## Arrancar Proyecto dev con Air
```
air init
air
```

## Ejecutar Test
```bash
go test ./... -v
```

## Docker alternativa

Para ejecutar dockerfile renombrar de .Dockerfile a Dockerfile ya que railway toma la configuracion de este.

```
docker build -t paynau-backend .
docker run -p 5000:5000 paynau-backend
```