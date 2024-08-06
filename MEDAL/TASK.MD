# MEDAL SERVICE

- Create Medal
- Update Medal
- Delete Medal

**Create Medal**: Service uses Message broker to retrieve data from the client, the process should be done in a transaction, a medal is created in this microservice, and then count of medals of the athlete is increased, then the transaction is committed, otherwise 
the transaction should be rolled back.

**Update Meda**: Service uses message broker, retrieves the data and does everything as `Create Medal` Service, which means the logic is done in a transaction.

**Delete Medal**: Service uses message broker, implements `soft-delete`, uses transaction to update the data inside the athlet's medals' count.
