
<div align="center">
<img src="https://drive.google.com/uc?id=1EcmbEcmvDJoPOH0_wHjY_poEuA52KMEP" height="60" />
</div>

# REAR

The **REAR protocol** facilitates the efficient exchange of resources and services through a standardized workflow:

- Providers advertise their available resources and services using a consistent format.
- Consumers search and discover resources based on specific criteria.
- REAR integrates with existing resource management systems and platforms.
- The protocol accommodates diverse resource types and allows for future expansions.

## Endpoints

- `GET /api/listflavours`: retrieve all available Flavours.
- `GET /api/listflavours/{flavourID}`: get a Flavour by the specified ID parameter.
- `GET /api/listflavours/selector`: list Flavours that match the provided selector.
- `POST /api/reserveflavour/{flavourID}`: reserve a Flavour specified by the ID, using the transaction details provided in the request body.
- `POST /api/purchaseflavour/{flavourID}`: purchase a Flavour specified by the ID, using the transaction details provided in the request body.
- `GET /api/listflavours/selector/syntax`: get the syntax for the selector.
- `GET /api/listflavours/selector/type`: get all available Flavour types.

These APIs allow users to interact with Flavours and transactions, enabling them to view, reserve, and purchase Flavours based on their needs.




