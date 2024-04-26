from sqlalchemy import Column, Integer, String
from sqlalchemy.ext.asyncio import (AsyncSession, async_sessionmaker,
                                    create_async_engine)
from sqlalchemy.ext.declarative import declarative_base

engine = create_async_engine("postgresql+asyncpg://user:password@postgres:5432/hackathone")

Base = declarative_base()

async_session = async_sessionmaker(engine, class_=AsyncSession, expire_on_commit=False)


async def get_db():
    async with async_session() as session:
        yield session
        
class Thread(Base):
    __tablename__ = "thread"
    id = Column(Integer, primary_key=True)
    image = Column(String, nullable=False)