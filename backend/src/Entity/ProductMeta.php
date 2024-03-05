<?php

namespace App\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Repository\ProductMetaRepository;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ProductMetaRepository::class)]
#[ApiResource]
class ProductMeta
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(type: Types::GUID)]
    private ?string $uuid = null;

    #[ORM\Column(length: 255)]
    private ?string $meta_key = null;

    #[ORM\Column(type: Types::TEXT, nullable: true)]
    private ?string $meta_value = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getUuid(): ?string
    {
        return $this->uuid;
    }

    public function setUuid(string $uuid): static
    {
        $this->uuid = $uuid;

        return $this;
    }

    public function getMetaKey(): ?string
    {
        return $this->meta_key;
    }

    public function setMetaKey(string $meta_key): static
    {
        $this->meta_key = $meta_key;

        return $this;
    }

    public function getMetaValue(): ?string
    {
        return $this->meta_value;
    }

    public function setMetaValue(?string $meta_value): static
    {
        $this->meta_value = $meta_value;

        return $this;
    }
}
