## omn3id

> Description: omn3id stand for omni incremental identifier
> Implementation Status: *V1Alpha*


spec:

- the idea behind omn3id is to be an unique incremental id provider, 
this is useful for backends that does not support incremental ids

why is needed:
- omniql depends that all their resources have incremental ids. 

know limitations:
- this can degrade the performance of the mutations in system that does not support incremental ids 


supported operations:

